package utils

import "testing"

func TestSum(t *testing.T) {

	type testdata struct {
		input  []int
		expect int
	}
	tests := []testdata{
		testdata{
			input:  []int{1, 2, 3, 4, 5, 6},
			expect: 21,
		},
		testdata{
			input:  []int{},
			expect: 0,
		},
		testdata{
			input:  []int{-1, -2, 2, 1},
			expect: 0,
		},
		testdata{
			input:  []int{1},
			expect: 1,
		},
	}
	for _, data := range tests {
		result := Sum(data.input...)
		if result != data.expect {
			t.Errorf("Expect %d ,but got %d", data.expect, result)
		}
	}

}
