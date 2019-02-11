package main

import (
	"fmt"
	"os"
	"text/template"
)

type person struct {
	Name string
	Age  int
}

func (p person) Info() string {
	return fmt.Sprintf("Name=%s,Age=%d", p.Name, p.Age)
}

var templateStr = `
Hi, {{.Name}}. I know you are {{ .Age}} years old.
`

var templateRangeStr = `
{{ range . }}
Hi, {{.Name}}. I know you are {{ .Age}} years old. (range version)
{{ end }}
`

var templateKVStr = `
{{ range $key, $value := . }}
Hi ID={{ $key }}, {{ $value.Name }}. I know you are {{ $value.Age }} years old. (map version)
{{ end }}
`

var templateKVWithMethodStr = `
{{ range $key, $value := . }}
Hi ID={{ $key }}, {{ $value.Info }}
{{ end }}
`

func main() {
	students := []person{
		{"Tom", 26},
		{"Bob", 24},
		{"Alice", 25},
		{"Mike", 29},
	}
	tmpl, err := template.New("test").Parse(templateStr)
	if err != nil {
		panic(err)
	}
	for _, s := range students {
		err = tmpl.Execute(os.Stdout, s)
		if err != nil {
			panic(err)
		}
	}
	// Hi, Tom. I know you are 26 years old.
	// Hi, Bob. I know you are 24 years old.
	// Hi, Alice. I know you are 25 years old.
	// Hi, Mike. I know you are 29 years old.

	rangeTmpl, err := template.New("rangetest").Parse(templateRangeStr)
	if err != nil {
		panic(err)
	}
	err = rangeTmpl.Execute(os.Stdout, students)
	if err != nil {
		panic(err)
	}
	// Hi, Tom. I know you are 26 years old. (range version)
	// Hi, Bob. I know you are 24 years old. (range version)
	// Hi, Alice. I know you are 25 years old. (range version)
	// Hi, Mike. I know you are 29 years old. (range version)

	studentsWithID := map[int]person{
		1: {"Tom", 26},
		2: {"Bob", 24},
		3: {"Alice", 25},
		4: {"Mike", 29},
	}
	kvtmpl, err := template.New("kvtest").Parse(templateKVStr)
	if err != nil {
		panic(err)
	}
	err = kvtmpl.Execute(os.Stdout, studentsWithID)
	if err != nil {
		panic(err)
	}
	// Hi ID=1, Tom. I know you are 26 years old. (map version)
	// Hi ID=2, Bob. I know you are 24 years old. (map version)
	// Hi ID=3, Alice. I know you are 25 years old. (map version)
	// Hi ID=4, Mike. I know you are 29 years old. (map version)

	kvMethodTmpl, err := template.New("kvMethodTest").Parse(templateKVWithMethodStr)
	if err != nil {
		panic(err)
	}
	err = kvMethodTmpl.Execute(os.Stdout, studentsWithID)
	if err != nil {
		panic(err)
	}
	// Hi ID=1, Name=Tom,Age=26
	// Hi ID=2, Name=Bob,Age=24
	// Hi ID=3, Name=Alice,Age=25
	// Hi ID=4, Name=Mike,Age=29
}
