package main

import (
	"os"
	"path/filepath"
	"time"

	"github.com/bamcop/kit"
	"github.com/bamcop/kit/debug"
	pf "github.com/bamcop/quickjs-react/pkg/polyfill"
	"github.com/buke/quickjs-go"
	polyfill "github.com/buke/quickjs-go-polyfill"
)

var (
	workingDir string
	srcCode    string
)

func init() {
	workingDir = debug.MustMainFileDir()

	filename := filepath.Join(workingDir, "../render_to_string/index.bundle_esbuild_cli.js")
	b, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	srcCode = string(b)
}

func main() {
	// Create a new runtime
	rt := quickjs.NewRuntime()
	defer rt.Close()

	// Create a new context
	ctx := rt.NewContext()
	defer ctx.Close()

	// Inject polyfills to the context
	polyfill.InjectAll(ctx)

	ret, err := ctx.Eval(pf.TryCatchWrap(srcCode))
	defer ret.Free()
	kit.Try(err)

	// Wait for the timeout to finish
	time.Sleep(time.Millisecond * 100)

	err = rt.ExecuteAllPendingJobs()
	kit.Try(err)
}
