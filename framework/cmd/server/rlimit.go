//+build ignore

package server

import (
	"github.com/c3sr/config"
	"github.com/c3sr/sys"
)

var (
	DefaultRLimitFileSoft uint64 = 500000
	DefaultRLimitFileHard uint64 = 500000
)

func setMaxFileRLimit() error {
	max := func(a, b uint64) uint64 {
		if a > b {
			return a
		}
		return b
	}

	softLimit := DefaultRLimitFileSoft
	hardLimit := DefaultRLimitFileHard

	sysSoftLimit, sysHardLimit, err := sys.GetMaxOpenFileLimit()
	if err == nil {
		softLimit = max(softLimit, sysSoftLimit)
		hardLimit = max(hardLimit, sysHardLimit)
	}

	return sys.SetMaxOpenFileLimit(softLimit, hardLimit)
}

func init() {
	config.AfterInit(func() {
		if err := setMaxFileRLimit(); err != nil {
			log.WithField("rlimit", err).Info("cannot set maximum file limit")
		}
	})
}
