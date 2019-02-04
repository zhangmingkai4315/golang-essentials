package benchmarking

import (
	"strings"
	"testing"
)

func BenchmarkCatString(b *testing.B) {
	strlist := []string{"hello", "world", "this", "is", "a", "test", "string", "array"}
	for i := 0; i < b.N; i++ {
		CatString(strlist, " ")
	}
}

func BenchmarkStringsJoin(b *testing.B) {
	strlist := []string{"hello", "world", "this", "is", "a", "test", "string", "array"}
	for i := 0; i < b.N; i++ {
		strings.Join(strlist, " ")
	}
}
