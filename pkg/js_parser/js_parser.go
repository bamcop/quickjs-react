package js_parser

import (
	"fmt"
	"strings"

	"github.com/samber/lo"
)

// ReplaceReactImport
// TODO ugly, need fix
func ReplaceReactImport(str string) string {
	if strings.Contains(str, "__toESM") {
		return str
	}
	pkg := strings.Split(str, `"`)[1]
	if pkg != "react" {
		return str
	}

	var buff string
	if strings.Contains(str, "{") {
		q := strings.Split(
			strings.Split(
				str,
				"{",
			)[1],
			"}",
		)[0]
		q = strings.ReplaceAll(q, " as", ":")

		buff += fmt.Sprintf("const {%s} = React;\n", q)
	}

	strs := strings.Split(str, " ")
	for _, s := range strs {
		if s == "{" {
			break
		}
		if !lo.Contains([]string{"import", "*", "as"}, s) {
			s = strings.TrimSuffix(s, ",")
			if s != "React" {
				buff += fmt.Sprintf("const %s = React;", s)
			}
			break
		}
	}

	return buff
}
