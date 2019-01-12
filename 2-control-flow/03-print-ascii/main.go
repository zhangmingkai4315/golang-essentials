// http://www.asciitable.com/

package main

import (
	"fmt"
)

func main() {
	for i := 33; i < 123; i++ {
		fmt.Printf("%v\t%x\t%#U \n", i, i, i)
	}
}

// 65      41      U+0041 'A'
// 66      42      U+0042 'B'
// 67      43      U+0043 'C'
// 68      44      U+0044 'D'
// 69      45      U+0045 'E'
// 70      46      U+0046 'F'
// 71      47      U+0047 'G'
