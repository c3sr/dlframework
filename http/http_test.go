package http

import (
	"os"
	"testing"

	"github.com/c3sr/config"
)

func TestMain(m *testing.M) {
	config.Init(
		config.AppName("carml"),
		config.VerboseMode(true),
		config.DebugMode(true),
	)
	os.Exit(m.Run())
}
