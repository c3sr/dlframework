// +build !nopython

package server

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	// "os/exec"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"time"

	sourcepath "github.com/GeertJohan/go-sourcepath"
	"github.com/c3sr/archive"
	dl "github.com/c3sr/dlframework"
	"github.com/c3sr/dlframework/framework/agent"
	dlcmd "github.com/c3sr/dlframework/framework/cmd"
	"github.com/c3sr/dlframework/framework/options"
	common "github.com/c3sr/dlframework/framework/predictor"
	"github.com/c3sr/dlframework/steps"
	"github.com/c3sr/go-python3"
	machine "github.com/c3sr/machine/info"
	nvidiasmi "github.com/c3sr/nvidia-smi"
	"github.com/c3sr/pipeline"
	"github.com/c3sr/tracer"
	"github.com/c3sr/uuid"
	"github.com/k0kubun/pp/v3"
	"github.com/levigross/grequests"
	opentracing "github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	jaeger "github.com/uber/jaeger-client-go"
	"github.com/unknwon/com"
)

var (
	inputsFilePath      string
	numInputParts       int
	numWarmUpInputParts int
)

func runPredictGeneralCmd(c *cobra.Command, args []string) error {
	if timeoutOptionSet {
		go func() {
			time.Sleep(15 * time.Minute)
			fmt.Println("timeout")
			os.Exit(-1)
		}()
	}

	// Initialize python interpreter
	python3.Py_Initialize()
	if !python3.Py_IsInitialized() {
		return errors.New("⚠️ Error initializing the python interpreter")
	}
	// See bugs and caveats https://docs.python.org/3/c-api/init.html#c.Py_FinalizeEx
	// defer python3.Py_Finalize()

	model, err := framework.FindModel(modelName + ":" + modelVersion)
	if err != nil {
		return err
	}
	log.WithField("model", modelName).Info("running predict inputs")

	predictors, err := agent.GetPredictors(framework)
	if err != nil {
		return errors.Wrapf(err,
			"⚠️ failed to get predictor for %s. make sure you have "+
				"imported the framework's predictor package",
			framework.MustCanonicalName(),
		)
	}

	var predictorHandle common.Predictor
	for _, pred := range predictors {
		predModality, err := pred.Modality()
		if err != nil {
			continue
		}
		modelModality, err := model.Modality()
		if err != nil {
			continue
		}
		if predModality == modelModality {
			predictorHandle = pred
			break
		}
	}
	if predictorHandle == nil {
		return errors.New("unable to find predictor for requested modality")
	}

	var dc map[string]int32
	if useGPU {
		if !nvidiasmi.HasGPU {
			return errors.New("not gpu found")
		}
		dc = map[string]int32{"GPU": 0}
		log.WithField("gpu = ", nvidiasmi.Info.GPUS[gpuDeviceId].ProductName).Info("Running evalaution on GPU")
	} else {
		dc = map[string]int32{"CPU": 0}
	}
	execOpts := &dl.ExecutionOptions{
		TraceLevel: dl.ExecutionOptions_TraceLevel(
			dl.ExecutionOptions_TraceLevel_value[traceLevel.String()],
		),
		DeviceCount: dc,
	}
	predOpts := &dl.PredictionOptions{
		FeatureLimit:     10,
		BatchSize:        int32(batchSize),
		GpuMetrics:       gpuMetrics,
		ExecutionOptions: execOpts,
	}

	rootSpan, ctx := tracer.StartSpanFromContext(
		context.Background(),
		tracer.APPLICATION_TRACE,
		"evaluation_predict_inputs",
		opentracing.Tags{
			"framework_name":     framework.Name,
			"framework_version":  framework.Version,
			"model_name":         modelName,
			"model_version":      modelVersion,
			"batch_size":         batchSize,
			"use_gpu":            useGPU,
			"gpu_metrics":        gpuMetrics,
			"num_warmup_batches": numWarmUpInputParts,
		},
	)
	if rootSpan == nil {
		panic("invalid span")
	}
	rootSpan.SetTag("num_batches", numInputParts)

	predictor, err := predictorHandle.Load(
		ctx,
		*model,
		options.PredictorOptions(predOpts),
		options.DisableFrameworkAutoTuning(disableFrameworkAutoTuning),
	)
	if err != nil {
		return err
	}

	var inputs []string
	inputsFilePath, err := filepath.Abs(inputsFilePath)
	if err != nil {
		return errors.Wrapf(err, "cannot get absolute path of %s", inputsFilePath)
	}

	f, err := os.Open(inputsFilePath)
	if err != nil {
		return errors.Wrapf(err, "cannot read %s", inputsFilePath)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		inputs = append(inputs, line)
	}

	log.WithField("inputs_file_path", inputsFilePath).
		Debug("using the specified inputs file path")

	if len(inputs) == 0 {
		log.WithError(err).Error("the inputs file has no input")
		os.Exit(-1)
	}

	tmp := inputs
	if duplicateInput < batchSize {
		duplicateInput = batchSize
	}
	for ii := 1; ii < duplicateInput; ii++ {
		inputs = append(inputs, tmp...)
	}

	_, modelManifest, err := predictor.Info()
	if err != nil {
		log.WithError(err).Error("can't find model manifest")
		os.Exit(-1)
	}

	// Not sure whether before preprocessed should be done here or just outside the program
	/*
	  beforePreprocess := strings.Split(modelManifest.GetBeforePreprocess(), " ")
	  if len(beforePreprocess) != 0 {
	    cmd := exec.Command(beforePreprocess[0], beforePreprocess[1:len(beforePreprocess)]...)
	    _, err := cmd.CombinedOutput()
	    if err != nil {
	      log.WithError(err).Error("before preprocess failed.")
	      os.Exit(-1)
	    }
	  }
	*/

	preprocessOptions, err := predictor.GetPreprocessOptions()

	inputParts := dl.PartitionStringList(inputs, partitionListSize)

	partlabels := map[string]string{}

	log.WithField("model", modelName).
		WithField("inputs_length", len(inputs)).
		WithField("input_parts_length", len(inputParts)).
		WithField("partition_list_size", partitionListSize).
		WithField("num_input_part", numInputParts).
		WithField("num_warmup_input_parts", numWarmUpInputParts).
		WithField("using_gpu", useGPU).
		Info("starting inference on inputs")

	pyState := python3.PyEval_SaveThread()

	if numWarmUpInputParts != 0 {
		warmUpSpan, warmUpSpanCtx := tracer.StartSpanFromContext(
			ctx,
			tracer.APPLICATION_TRACE,
			"warm_up",
			opentracing.Tags{
				"num_warmup_batches": numWarmUpInputParts,
			},
		)

		tracer.SetLevel(tracer.NO_TRACE)

		for _, part := range inputParts[0:numWarmUpInputParts] {
			input := make(chan interface{}, DefaultChannelBuffer)
			go func() {
				defer close(input)
				for _, inp := range part {
					id := uuid.NewV4()
					lbl := steps.NewIDWrapper(id, inp)
					partlabels[lbl.GetID()] = "" // no label for the input
					input <- lbl
				}
			}()

			opts := []pipeline.Option{pipeline.ChannelBuffer(DefaultChannelBuffer)}
			if tracePreprocess == true {
				opts = append(opts, pipeline.Context(warmUpSpanCtx))
			}
			output := pipeline.New(opts...).
				Then(steps.NewPreprocessGeneral(preprocessOptions, modelManifest.GetPreprocess())).
				Run(input)

			var tensors []interface{}
			for out := range output {
				if err, ok := out.(error); ok && failOnFirstError {
					log.WithError(err).Error("encountered an error while performing preprocessing")
					os.Exit(-1)
				}
				tensors = append(tensors, out)
			}

			tensorParts := dl.Partition(tensors, batchSize)

			input = make(chan interface{}, DefaultChannelBuffer)
			go func() {
				defer close(input)
				for _, p := range tensorParts {
					input <- p
				}
			}()

			output = pipeline.New(pipeline.Context(warmUpSpanCtx), pipeline.ChannelBuffer(DefaultChannelBuffer)).
				Then(steps.NewPredictGeneral(predictor, modelManifest.GetPostprocess())).
				Run(input)

			for o := range output {
				if err, ok := o.(error); ok && failOnFirstError {
					log.WithError(err).Error("encountered an error while performing warmup inference")
					os.Exit(-1)
				}
			}
		}

		tracer.SetLevel(traceLevel)

		warmUpSpan.Finish()
	}

	outputs := make(chan interface{}, DefaultChannelBuffer)

	if numInputParts == -1 {
		numInputParts = len(inputParts)
	}

	hostName, _ := os.Hostname()
	metadata := map[string]string{}
	if useGPU {
		if bts, err := json.Marshal(nvidiasmi.Info); err == nil {
			metadata["nvidia_smi"] = string(bts)
			rootSpan.SetTag("nvidia_smi", string(bts))
		}
	}

	inputCnt := len(inputs)
	inferenceProgress := dlcmd.NewProgress("inferring", inputCnt)

	for ii, part := range inputParts[:numInputParts] {
		input := make(chan interface{}, DefaultChannelBuffer)
		go func() {
			defer close(input)
			for _, inp := range part {
				id := uuid.NewV4()
				lbl := steps.NewIDWrapper(id, inp)
				partlabels[lbl.GetID()] = "" // no label for the input
				input <- lbl
			}
		}()

		evaluateBatchSpan, evaluateBatchCtx := tracer.StartSpanFromContext(
			ctx,
			tracer.APPLICATION_TRACE,
			"evaluate_batch",
			opentracing.Tags{
				"batch_index": ii,
			},
		)

		opts := []pipeline.Option{pipeline.ChannelBuffer(DefaultChannelBuffer)}
		if tracePreprocess == true {
			opts = append(opts, pipeline.Context(evaluateBatchCtx))
		}
		output := pipeline.New(opts...).
			Then(steps.NewPreprocessGeneral(preprocessOptions, modelManifest.GetPreprocess())).
			Run(input)

		var tensors []interface{}
		for out := range output {
			if err, ok := out.(error); ok && failOnFirstError {
				inferenceProgress.Finish()

				log.WithError(err).Error("encountered an error while performing preprocessing")
				os.Exit(-1)
			}
			tensors = append(tensors, out)
		}

		tensorParts := dl.Partition(tensors, batchSize)

		input = make(chan interface{}, DefaultChannelBuffer)
		go func() {
			defer close(input)
			for _, p := range tensorParts {
				input <- p
			}
		}()

		output = pipeline.New(pipeline.Context(evaluateBatchCtx), pipeline.ChannelBuffer(DefaultChannelBuffer)).
			Then(steps.NewPredictGeneral(predictor, modelManifest.GetPostprocess())).
			Run(input)

		inferenceProgress.Add(batchSize)

		for o := range output {
			if err, ok := o.(error); ok && failOnFirstError {
				inferenceProgress.Finish()

				log.WithError(err).Error("encountered an error while performing inference")
				os.Exit(-1)
			}
			if saveInferenceResult {
				outputs <- o
			}
		}
		evaluateBatchSpan.Finish()
	}

	inferenceProgress.Finish()

	predictor.Close()

	rootSpan.Finish()
	tracer.ResetStd()

	close(outputs)

	python3.PyEval_RestoreThread(pyState)

	traceID := rootSpan.Context().(jaeger.SpanContext).TraceID()
	traceIDVal := traceID.String()
	if runtime.GOARCH == "ppc64le" {
		traceIDVal = strconv.FormatUint(traceID.Low, 16)
	}
	tracerServerAddr := getTracerServerAddress(tracerAddress)
	pp.Println(fmt.Sprintf("the trace is at http://%s:16686/trace/%v", tracerServerAddr, traceIDVal))
	traceURL := fmt.Sprintf("http://%s:16686/api/traces/%v?raw=true", tracerServerAddr, traceIDVal)

	var device string
	if useGPU {
		device = "gpu"
	} else {
		device = "cpu"
		gpuDeviceId = -1
	}

	if publishToDatabase == false {

		outputDir := filepath.Join(baseDir, framework.Name, framework.Version, model.Name, model.Version, strconv.Itoa(batchSize), device, hostName)
		if !com.IsDir(outputDir) {
			os.MkdirAll(outputDir, os.ModePerm)
		}

		if useGPU {
			if bts, err := json.Marshal(nvidiasmi.Info); err == nil {
				ioutil.WriteFile(filepath.Join(outputDir, "nvidia_info.json"), bts, 0644)
			}
		}

		if machine.Info != nil && machine.Info.Hostname != "" {
			bts, err := json.Marshal(machine.Info)
			if err == nil {
				ioutil.WriteFile(filepath.Join(outputDir, "system_info.json"), bts, 0644)
			}
		}

		ts := strings.ToLower(tracer.LevelToName(traceLevel))
		traceFileName := "trace_" + ts + ".json"
		tracePath := filepath.Join(outputDir, traceFileName)
		if (publishToDatabase == false) && com.IsFile(tracePath) {
			log.WithField("path", tracePath).Info("trace file already exists")
			return nil
		}

		resp, err := grequests.Get(traceURL, nil)
		if err != nil {
			log.WithError(err).
				WithField("trace_id", traceIDVal).
				Error("failed to download span information")
		}
		log.WithField("model", modelName).WithField("trace_id", traceIDVal).WithField("traceURL", traceURL).Info("downloaded trace information")

		err = ioutil.WriteFile(tracePath, resp.Bytes(), 0644)
		if err != nil {
			return err
		}

		if false {
			archiver, err := archive.Zip(tracePath, archive.BZip2Format())
			if err != nil {
				return err
			}

			archiveFilePath := strings.TrimSuffix(tracePath, ".json") + ".tar.bz2"
			f, err := os.Create(archiveFilePath)
			if err != nil {
				return err
			}
			defer f.Close()

			_, err = io.Copy(f, archiver)
		}

		log.WithField("model", modelName).WithField("path", tracePath).Infof("publishToDatabase is false, writing the trace to a local file")

		pp.Println(fmt.Sprintf("the trace is at %v locally", tracePath))

		return nil
	}

	return nil
}

var predictGeneralCmd = &cobra.Command{
	Use:     "general",
	Short:   "Evaluate using general methods",
	Aliases: []string{"generals"},
	RunE: func(c *cobra.Command, args []string) error {
		if modelName == "all" {
			for _, model := range framework.Models() {
				modelName = model.Name
				modelVersion = model.Version
				runPredictGeneralCmd(c, args)
			}
			return nil
		}
		return runPredictGeneralCmd(c, args)
	},
}

func init() {
	sourcePath := sourcepath.MustAbsoluteDir()
	defaultInputsPath := filepath.Join(sourcePath, "..", "_fixtures", "urlsfile")
	if !com.IsFile(defaultInputsPath) {
		defaultInputsPath = ""
	}
	predictGeneralCmd.PersistentFlags().StringVar(&inputsFilePath, "inputs_file_path", defaultInputsPath, "the path of the file containing the inputs to perform the evaluations on.")
	predictGeneralCmd.PersistentFlags().IntVar(&numInputParts, "num_input_parts", -1, "the number of input parts to process. Setting input parts to a value other than -1 means that only the first num_input_parts * partition_list_size tensors are infered from the dataset. This is useful while performing performance evaluations, where only a few hundred evaluation samples are useful")
	predictGeneralCmd.PersistentFlags().IntVar(&numWarmUpInputParts, "num_warmup_input_parts", 1, "the number of input parts to process during the warmup period. This is ignored if num_file_parts=-1")
}
