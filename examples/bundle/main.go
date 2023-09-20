package main

import (
	"os"

	"github.com/bamcop/kit/debug"
	"github.com/evanw/esbuild/pkg/api"
)

func main() {
	result := api.Build(api.BuildOptions{
		AbsWorkingDir: debug.MustMainFileDir(),
		EntryPoints:   []string{"input.js"},
		Outfile:       "output.js",
		Bundle:        true,
		Write:         true,
		LogLevel:      api.LogLevelInfo,
	})

	if len(result.Errors) > 0 {
		os.Exit(1)
	}
}
