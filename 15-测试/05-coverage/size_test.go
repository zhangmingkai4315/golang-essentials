package size

import "testing"

func TestSize(t *testing.T) {
	tests := []struct {
		in  int
		out string
	}{
		{
			in:  0,
			out: "zero",
		},
		{
			in:  5,
			out: "small",
		},
	}
	for _, test := range tests {
		size := Size(test.in)
		if size != test.out {
			t.Errorf("Size(%d) = %s; expect %s", test.in, size, test.out)
		}
	}
}
