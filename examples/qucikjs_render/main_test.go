package main

import (
	"testing"
)

func BenchmarkEvalRenderToString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		EvalRenderToString()
	}
}
