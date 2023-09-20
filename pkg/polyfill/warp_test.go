package polyfill

import (
	"testing"

	"github.com/buke/quickjs-go"
	polyfill "github.com/buke/quickjs-go-polyfill"
	"github.com/stretchr/testify/assert"
)

func TestTryCatchWrap(t *testing.T) {
	rt := quickjs.NewRuntime()
	defer rt.Close()

	ctx := rt.NewContext()
	defer ctx.Close()

	// Inject polyfills to the context
	polyfill.InjectAll(ctx)

	ret1, err1 := ctx.Eval(`throw 'abc'`)
	defer ret1.Free()
	assert.Nil(t, err1)

	ret2, err2 := ctx.Eval(TryCatchWrap(`throw 'abc'`))
	defer ret2.Free()
	assert.NotNil(t, err2)
}
