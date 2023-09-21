package main

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/bamcop/kit"
	"github.com/bamcop/kit/debug"
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
		b, err := os.ReadFile(filepath.Join(RootDir, "public/examples/hello.html"))
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		c.Writer.Header().Set("Content-Type", `text/html; charset=UTF-8`)
		_, _ = c.Writer.WriteString(string(b))
		c.Status(http.StatusOK)
	})

	r.Run(":4399")
}
