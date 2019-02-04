package utils

import "testing"

func TestSum(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6}
	expect := 21
	result := Sum(input...)
	if result != expect {
		t.Errorf("Expect %d ,but got %d", expect, result)
	}
}
