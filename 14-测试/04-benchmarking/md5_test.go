package benchmarking

import (
	"fmt"
	"testing"
)

func TestGetStringMd5(t *testing.T) {
	input := "hello world"
	expect := "5eb63bbbe01eeed093cb22bb8f5acdc3"
	result := GetStringMd5(input)
	if result != expect {
		t.Errorf("expect %s but got %s", expect, result)
	}
}

func ExampleGetStringMd5() {
	input := "hello world"
	fmt.Println(GetStringMd5(input))
	// Output
	// 5eb63bbbe01eeed093cb22bb8f5acdc3
}
func BenchmarkGetStringMd5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetStringMd5("hello world")
	}
}
