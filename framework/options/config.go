package options

import (
	"github.com/k0kubun/pp/v3"
	"github.com/c3sr/config"
	"github.com/c3sr/vipertags"
)

type optionsConfig struct {
	BatchSize           int    `json:"batch_size" config:"predictor.batch_size"`
	GPUMetrics          string `json:"gpu_metrics" config:"predictor.gpu_metrics"`
	FeatureLimit        int    `json:"feature_limit" config:"predictor.batch_size"`
	DefaultDeviceString string `json:"default_device" config:"predictor.default_device"`
	DefaultDevice       device `json:"-" config:"-"`
	done                chan struct{}
}

var (
	Config = &optionsConfig{
		done: make(chan struct{}),
	}
)

func (optionsConfig) ConfigName() string {
	return "predictor/options"
}

func (a *optionsConfig) SetDefaults() {
	vipertags.SetDefaults(a)
}

func (a *optionsConfig) Read() {
	defer close(a.done)
	vipertags.Fill(a)
	if a.BatchSize == 0 {
		a.BatchSize = DefaultBatchSize
	}
	if a.FeatureLimit == 0 {
		a.FeatureLimit = DefaultFeatureLimit
	}
	if a.GPUMetrics == "" {
		a.GPUMetrics = DefaultGPUMetrics
	}
	if a.DefaultDeviceString == "" {
		a.DefaultDevice = DefaultDevice
	} else if a.DefaultDeviceString == "cuda" {
		a.DefaultDevice = device{deviceType: CUDA_DEVICE, id: 0}
	} else {
		a.DefaultDevice = device{deviceType: CPU_DEVICE, id: 0}
	}
}

func (c optionsConfig) Wait() {
	<-c.done
}

func (c optionsConfig) String() string {
	return pp.Sprintln(c)
}

func (c optionsConfig) Debug() {
	log.Debug("predictor/options Config = ", c)
}

func init() {
	config.Register(Config)
}
