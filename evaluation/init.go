package evaluation

import (
	"github.com/sirupsen/logrus"

	"github.com/c3sr/config"
	"github.com/c3sr/logger"
)

type logWrapper struct {
	*logrus.Entry
}

var (
	log = &logWrapper{
		Entry: logger.New().WithField("pkg", "evaluation"),
	}
)

func (l *logWrapper) Output(calldepth int, s string) error {
	// l.WithField("calldepth", calldepth).Debug(s)
	return nil
}

func init() {
	config.AfterInit(func() {
		log = &logWrapper{
			Entry: logger.New().WithField("pkg", "evaluation"),
		}
	})
}
