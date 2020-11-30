package cmd

import (
	"github.com/pkg/errors"
	"github.com/c3sr/nvidia-smi"
)

func NVIDIASmi() (*nvidiasmi.NvidiaSmi, error) {
	if !nvidiasmi.HasGPU {
		return nil, errors.New("no gpus found")
	}
	return nvidiasmi.Info, nil
}
