package main

import (
	"fmt"
	"net/http"
)

type Counter struct {
	n int
}

func (c *Counter) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c.n++
	fmt.Fprintf(w, "Counter=%d\n", c.n)
}

func main() {
	c := new(Counter)
	http.Handle("/counter", c)
	http.ListenAndServe("127.0.0.1:8080", nil)
}
