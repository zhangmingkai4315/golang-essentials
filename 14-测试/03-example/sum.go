// package example include some basic math function
package example

// Sum will receive unlimit number of int
// and return sum of all numbers
func Sum(arr ...int) (result int) {
	for _, i := range arr {
		result += i
	}
	return
}
