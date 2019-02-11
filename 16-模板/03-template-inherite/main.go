package main

import (
	"fmt"
	"html/template"
	"math/rand"
	"net/http"
	"time"
)

func process(w http.ResponseWriter, r *http.Request) {
	rand.Seed(time.Now().Unix())
	var t *template.Template
	var err error
	if rand.Intn(10) > 5 {
		t, err = template.ParseFiles("layout/base.tmpl", "index_blue.tmpl")
		if err != nil {
			fmt.Fprint(w, err)
			return
		}
	} else {
		t, err = template.ParseFiles("layout/base.tmpl", "index_red.tmpl")
		if err != nil {
			fmt.Fprint(w, err)
			return
		}
	}
	t.ExecuteTemplate(w, "layout", "")
}

func main() {
	http.HandleFunc("/", process)
	http.ListenAndServe(":8080", nil)
}
