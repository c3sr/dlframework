package cmd

import (
	"path/filepath"

	"github.com/k0kubun/pp/v3"

	"github.com/unknwon/com"
	"github.com/fatih/color"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/c3sr/config"
	"github.com/c3sr/logger"
	"github.com/sirupsen/logrus"
)

var (
	IsDebug   bool
	IsVerbose bool
	AppSecret string
	CfgFile   string
	log       *logrus.Entry = logrus.New().WithField("pkg", "dlframework/framework/cmd")
)

// Init reads in config file and ENV variables if set.
func Init() {

	log.Level = logrus.DebugLevel
	config.AfterInit(func() {
		log = logger.New().WithField("pkg", "dlframework/framework/cmd")
	})

	color.NoColor = false
	opts := []config.Option{
		config.AppName("carml"),
		config.ColorMode(true),
		config.DebugMode(IsDebug),
		config.VerboseMode(IsVerbose),
	}
	if IsDebug || IsVerbose {
		pp.WithLineInfo = true
	}
	if c, err := homedir.Expand(CfgFile); err == nil {
		CfgFile = c
	}
	if config.IsValidRemotePrefix(CfgFile) {
		opts = append(opts, config.ConfigRemotePath(CfgFile))
	} else if com.IsFile(CfgFile) {
		if c, err := filepath.Abs(CfgFile); err == nil {
			CfgFile = c
		}
		opts = append(opts, config.ConfigFileAbsolutePath(CfgFile))
	}

	if AppSecret != "" {
		opts = append(opts, config.AppSecret(AppSecret))
	}
	config.Init(opts...)
}
