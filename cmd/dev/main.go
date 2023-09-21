package main

import (
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

	r.Run(":4399")
}
