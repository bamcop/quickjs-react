package main

import (
	"os"
	"path/filepath"

	"github.com/bamcop/kit"
	"github.com/bamcop/kit/debug"
	"github.com/evanw/esbuild/pkg/api"
)

func main() {
	workingDir, err := filepath.Abs(filepath.Join(debug.MustMainFileDir(), "../.."))
	kit.Try(err)

	result := api.Build(api.BuildOptions{
		AbsWorkingDir: workingDir,
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
