package agent

import (
	"encoding/json"
	"os"
	"path"
	"runtime"
	"strings"
	"sync"

	cpuinfo "github.com/c3sr/machine/info"
	nvidiasmi "github.com/c3sr/nvidia-smi"

	"github.com/c3sr/config"
	dl "github.com/c3sr/dlframework"
	store "github.com/c3sr/libkv/store"
	"github.com/c3sr/registry"
)

type base struct {
	Framework dl.FrameworkManifest
}

func toPath(s string) string {
	return strings.Replace(s, ":", "/", -1)
}

func (b *base) PublishInPredictor(host, prefix string) error {

	ttl := registry.Config.Timeout
	marshaler := registry.Config.Serializer

	rgs, err := registry.New()
	if err != nil {
		return err
	}
	defer rgs.Close()

	prefix = path.Join(config.App.Name, prefix)

	frameworks0, err := dl.Frameworks()
	if err != nil {
		log.WithError(err).Error("failed to get frameworks while publishing predictor")
		return err
	}
	frameworks := make([]*dl.FrameworkManifest, len(frameworks0))
	for ii := range frameworks0 {
		frameworks[ii] = &frameworks0[ii]
	}

	var wg sync.WaitGroup
	models := b.Framework.Models()
	wg.Add(len(models))

	for _, model := range models {
		func(model dl.ModelManifest) {
			defer wg.Done()
			mn, err := model.CanonicalName()
			if err != nil {
				return
			}
			spltHost := strings.Split(host, ":")
			ip, port := spltHost[0], spltHost[1]
			hostName, _ := os.Hostname()
			gpuinfo, err := json.Marshal(nvidiasmi.Info)
			if err != nil {
				log.WithError(err).Error("failed to get agent's nvidia-smi information")
				gpuinfo = []byte{}
			}
			cpuinfo, err := json.Marshal(cpuinfo.Info)
			if err != nil {
				log.WithError(err).Error("failed to get agent's cpu information")
				gpuinfo = []byte{}
			}
			bts, err := marshaler.Marshal(&dl.Agent{
				Host:         ip,
				Port:         port,
				Hostname:     hostName,
				Architecture: runtime.GOARCH,
				Hasgpu:       nvidiasmi.HasGPU,
				Gpuinfo:      string(gpuinfo),
				Cpuinfo:      string(cpuinfo),
				Frameworks:   frameworks,
			})
			if err != nil {
				return
			}

			key := path.Join(prefix, toPath(mn), "agent-"+host)
			rgs.Put(key, bts, &store.WriteOptions{TTL: ttl, IsDir: false})
		}(model)
	}

	wg.Wait()

	return nil
}

func (b *base) PublishInRegistery(prefix string) error {

	framework := b.Framework
	cn, err := framework.CanonicalName()
	if err != nil {
		return err
	}

	ttl := registry.Config.Timeout
	marshaler := registry.Config.Serializer

	rgs, err := registry.New()
	if err != nil {
		return err
	}
	defer rgs.Close()

	prefix = path.Join(config.App.Name, prefix)

	key := path.Join(prefix, "frameworks")

	key = path.Join(key, cn)

	rgs.Put(key, []byte(cn), &store.WriteOptions{TTL: ttl, IsDir: false})

	key = path.Join(prefix, toPath(cn))

	key = path.Join(key, "manifest.json")
	bts, err := marshaler.Marshal(&framework)
	if err != nil {
		return err
	}

	rgs.Put(key, bts, &store.WriteOptions{TTL: ttl, IsDir: false})

	var wg sync.WaitGroup
	models := framework.Models()
	wg.Add(len(models))

	for _, model := range models {
		func(model dl.ModelManifest) {
			defer wg.Done()
			mn, err := model.CanonicalName()
			if err != nil {
				return
			}
			bts, err := marshaler.Marshal(&model)
			if err != nil {
				return
			}
			key := path.Join(prefix, toPath(mn), "manifest.json")
			rgs.Put(key, bts, &store.WriteOptions{TTL: ttl, IsDir: false})
		}(model)
	}

	wg.Wait()

	return nil
}
