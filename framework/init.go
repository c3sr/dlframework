package framework

import (
	"github.com/c3sr/config"
	"github.com/c3sr/logger"
	"github.com/sirupsen/logrus"
)

var (
	log       *logrus.Entry = logrus.New().WithField("pkg", "dlframework/framework")
	debugging               = false
)

func init() {
	log.Level = logrus.DebugLevel
	config.AfterInit(func() {
		log = logger.New().WithField("pkg", "dlframework/framework")
	})
}
