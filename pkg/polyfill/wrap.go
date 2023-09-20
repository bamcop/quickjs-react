package polyfill

import (
	"strings"
)

const WrapCode = `
try {
  // TODO
} catch (error) {
  console.error('try catch wrap:', error)
  error.___raise___()
}
`

// TryCatchWrap
// current, ctx.Eval(`throw 'abc'`) return err == nil, add TryCatchWrap for catch exception
func TryCatchWrap(code string) string {
	return strings.ReplaceAll(WrapCode, "// TODO", code)
}
