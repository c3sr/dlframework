// +build ignore

package main

import (
	"fmt"
	"os"

	"github.com/c3sr/config"
	"github.com/c3sr/dlframework/framework/cmd"
	evalcmd "github.com/c3sr/dlframework/evaluation/cmd"
	"github.com/sirupsen/logrus"
)

var (
	log *logrus.Entry = logrus.New().WithField("pkg", "dlframework/framework/cmd/evaluate")
)

func main() {
	config.AfterInit(func() {
		log = logrus.New().WithField("pkg", "dlframework/framework/cmd/evaluate")
	})

	cmd.IsDebug = true
	cmd.IsVerbose = true
	cmd.Init()

	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logrus.DebugLevel)
	if err := evalcmd.EvaluationCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
