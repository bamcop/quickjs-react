package debug

import (
	"bytes"
	"debug/gosym"
	"debug/macho"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/bamcop/kit"
	"github.com/goretk/gore"
)

// MainFilePath Unable to get the location of main.go through debug.Stack() in init function
// src: vulncheck/internal/buildinfo/buildinfo.go: openExe
// src: vulncheck/internal/buildinfo/additions_scan.go: ExtractPackagesAndSymbols
// src: debug/gosym/pclntab_test.go: TestPCLine
func MainFilePath() kit.Result[string] {
	// github.com/goretk/gore@v0.11.1 not supported Arm Darwin
	// TODO: not real main.main filepath
	if runtime.GOOS != "darwin" {
		f, err := gore.Open(os.Args[0])
		if err != nil {
			return kit.NewResultE[string](err)
		}
		pkgs, err := f.GetPackages()
		if err != nil {
			return kit.NewResultE[string](err)
		}
		for _, pkg := range pkgs {
			if pkg.Name == "main" {
				return kit.NewResultV(filepath.Join(pkg.Filepath, "main.go"))
			}
		}

		return kit.NewResultE[string](errors.New("not found main package"))
	}

	b, err := os.ReadFile(os.Args[0])
	if err != nil {
		return kit.NewResultE[string](err)
	}

	f, err := macho.NewFile(bytes.NewReader(b))
	if err != nil {
		return kit.NewResultE[string](err)
	}

	var textOffset uint64
	text := f.Section("__text")
	if text != nil {
		textOffset = uint64(text.Offset)
	}

	pclntab := f.Section("__gopclntab")
	if pclntab == nil {
		return kit.NewResultE[string](fmt.Errorf("gopclntab is nil"))
	}
	pclndat, err := pclntab.Data()
	if err != nil {
		return kit.NewResultE[string](err)
	}

	lineTab := gosym.NewLineTable(pclndat, textOffset)
	tab, err := gosym.NewTable(nil, lineTab)
	if err != nil {
		return kit.NewResultE[string](err)
	}

	fn := tab.LookupFunc("main.main")
	file, _, _ := tab.PCToLine(fn.Sym.Value)

	return kit.NewResultV(file)
}

func MustMainFileDir() string {
	return filepath.Dir(MainFilePath().Unwrap())
}
