package main

import (
	"os"
	"path/filepath"

	"github.com/bamcop/kit"
	"github.com/bamcop/kit/debug"
	"github.com/evanw/esbuild/pkg/api"
)

func main() {
	absWorkingDir, err := filepath.Abs(filepath.Join(debug.MustMainFileDir(), "../render_to_string"))
	kit.Try(err)

	result := api.Build(api.BuildOptions{
		AbsWorkingDir: absWorkingDir,
		EntryPoints:   []string{"index.js"},
		Outfile:       "index.bundle_esbuild_go_api.js",
		Bundle:        true,
		Write:         true,
		LogLevel:      api.LogLevelInfo,
		Platform:      api.PlatformNode,
		Loader:        map[string]api.Loader{".js": api.LoaderJSX},
	})

	if len(result.Errors) > 0 {
		os.Exit(1)
	}
}
