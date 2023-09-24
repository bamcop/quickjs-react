package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/bamcop/kit"
	"github.com/bamcop/kit/debug"
	"github.com/bamcop/quickjs-react/pkg/js_parser"
	pf "github.com/bamcop/quickjs-react/pkg/polyfill"
	"github.com/bamcop/quickjs-react/pkg/polyfill/text_encoder"
	"github.com/buke/quickjs-go"
	polyfill "github.com/buke/quickjs-go-polyfill"
	"github.com/evanw/esbuild/pkg/api"
	"github.com/gin-gonic/gin"
)

var (
	RootDir string
)

func init() {
	var err error
	RootDir, err = filepath.Abs(filepath.Join(debug.MustMainFileDir(), "../.."))
	kit.Try(err)
}

func main() {
	r := gin.Default()

	r.Static("/static", filepath.Join(RootDir, "public"))
	r.StaticFile("/favicon.ico", filepath.Join(RootDir, "public", "favicon.ico"))

	r.GET("/", func(c *gin.Context) {
		b, err := BundleServer()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err,
			})
			return
		}

		str, err := RenderAppToString(b)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err,
			})
			return
		}

		c.Writer.Header().Set("Content-Type", `text/html; charset=UTF-8`)
		_, _ = c.Writer.WriteString(str)
		c.Status(http.StatusOK)
	})

	r.GET("/app.client.js", func(c *gin.Context) {
		b, err := BundleClient()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err,
			})
			return
		}
		//os.WriteFile(filepath.Join(RootDir, "tmp", "app.client.js"), b, 0644)

		var (
			code    = string(b)
			reg     = regexp.MustCompile(`import.*;`)
			matches = reg.FindAllStringSubmatch(code, -1)
		)

		for _, olds := range matches {
			news := js_parser.ReplaceReactImport(olds[0])
			code = strings.ReplaceAll(code, olds[0], news)
		}

		code = strings.ReplaceAll(code, `__require("react")`, `React`)

		c.Writer.Header().Set("Content-Type", `application/javascript`)
		_, _ = c.Writer.WriteString(code)
		c.Status(http.StatusOK)
	})

	r.Run(":4399")
}

func BundleServer() ([]byte, error) {
	result := api.Build(api.BuildOptions{
		AbsWorkingDir: RootDir,
		EntryPoints:   []string{"src/index.js"},
		Bundle:        true,
		Format:        api.FormatIIFE,
		Write:         false,
		LogLevel:      api.LogLevelInfo,
		Platform:      api.PlatformDefault,
		Loader:        map[string]api.Loader{".js": api.LoaderJSX},
	})

	if len(result.Errors) > 0 {
		slog.Error("esbuild", slog.Any("err", result.Errors))
		return nil, fmt.Errorf("%v", result.Errors)
	}

	return result.OutputFiles[0].Contents, nil
}

func BundleClient() ([]byte, error) {
	result := api.Build(api.BuildOptions{
		AbsWorkingDir: RootDir,
		EntryPoints:   []string{"src/app.jsx"},
		Bundle:        true,
		Format:        api.FormatESModule,
		Write:         false,
		LogLevel:      api.LogLevelInfo,
		Platform:      api.PlatformDefault,
		Loader:        map[string]api.Loader{".js": api.LoaderJSX},
		External:      []string{"react"},
	})

	if len(result.Errors) > 0 {
		slog.Error("esbuild", slog.Any("err", result.Errors))
		return nil, fmt.Errorf("%v", result.Errors)
	}

	return result.OutputFiles[0].Contents, nil
}

func RenderAppToString(b []byte) (string, error) {
	// Create a new runtime
	rt := quickjs.NewRuntime()
	defer rt.Close()

	// Create a new context
	ctx := rt.NewContext()
	defer ctx.Close()

	// Inject polyfills to the context
	polyfill.InjectAll(ctx)

	// customer polyfill
	{
		err := text_encoder.InjectTo(ctx)
		kit.Try(err)
	}

	ret, err := ctx.Eval(pf.TryCatchWrap(string(b)))
	defer ret.Free()
	kit.Try(err)

	ret = ctx.Globals().Call("RenderApp")
	defer ret.Free()

	htmlString := wrapHtml(ret.String())
	return htmlString, nil
}

func wrapHtml(src string) string {
	b, err := os.ReadFile(filepath.Join(RootDir, "src/index.html"))
	kit.Try(err)

	str := strings.ReplaceAll(string(b), "<h1>Hello Golang</h1>", src)
	return str
}
