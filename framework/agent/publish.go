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
	lock "github.com/c3sr/lock/registry"
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

	rgs.Put(prefix, nil, &store.WriteOptions{IsDir: true})

	var wg sync.WaitGroup
	models := b.Framework.Models()
	wg.Add(len(models))
	for _, model := range models {
		go func(model dl.ModelManifest) {
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

	rgs.Put(prefix, nil, &store.WriteOptions{IsDir: true})

	locker := lock.New(rgs)
	locker.Lock(prefix)
	defer locker.Unlock(prefix)

	frameworksKey := path.Join(prefix, "frameworks")

	kv, err := rgs.Get(frameworksKey)
	if err != nil {
		if ok, e := rgs.Exists(frameworksKey); e == nil && ok {
			log.WithError(err).Errorf("cannot get value for key %v", frameworksKey)
			return err
		}
		kv = &store.KVPair{
			Key:   frameworksKey,
			Value: []byte{},
		}
	}
	found := false
	val := strings.TrimSpace(string(kv.Value))
	frameworkLines := strings.Split(val, "\n")
	for _, name := range frameworkLines {
		if name == cn {
			found = true
			break
		}
	}
	if !found {
		frameworkLines = append(frameworkLines, cn)
		newVal := strings.TrimSpace(strings.Join(frameworkLines, "\n"))
		rgs.AtomicPut(frameworksKey, []byte(newVal), kv, nil)
	}

	key := path.Join(prefix, toPath(cn))
	rgs.Put(key, nil, &store.WriteOptions{TTL: ttl, IsDir: true})

	key = path.Join(key, "manifest.json")
	bts, err := marshaler.Marshal(&framework)
	if err != nil {
		return err
	}
	if err := rgs.Put(key, bts, &store.WriteOptions{TTL: ttl, IsDir: false}); err != nil {
		return err
	}

	var wg sync.WaitGroup
	models := framework.Models()
	wg.Add(len(models))
	for _, model := range models {
		go func(model dl.ModelManifest) {
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
