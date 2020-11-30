package client

import (
	"github.com/c3sr/config"
	"github.com/c3sr/logger"
	_ "github.com/c3sr/tracer/jaeger"
	_ "github.com/c3sr/tracer/noop"
	"github.com/sirupsen/logrus"
)

var (
	log *logrus.Entry
)

func init() {
	config.AfterInit(func() {
		log = logger.New().WithField("pkg", "dlframework/framework/cmd/client")
	})

}
