package text_encoder

import (
	_ "embed"

	"github.com/buke/quickjs-go"
)

//go:embed js/text_encoder.js
var encoderJs string

func InjectTo(ctx *quickjs.Context) error {
	ret, err := ctx.Eval(encoderJs)
	defer ret.Free()

	if err != nil {
		panic(err)
	}

	return nil
}
