package main

import (
	"os"
	"text/template"
)

type user struct {
	Name     string
	Phones   []string
	IsAdmin  bool
	IsLoggin bool
}

func main() {
	mike := user{"Mike", []string{"1234567890", "010-123456"}, true, true}
	indexTplStr := `
the first phone number is {{index .Phones 0 }}
`
	indexFuncTpl, err := template.New("indextest").Parse(indexTplStr)
	if err != nil {
		panic(err)
	}
	err = indexFuncTpl.Execute(os.Stdout, mike)
	if err != nil {
		panic(err)
	}
	// the first phone number is 1234567890

	andTplStr := `
the user is{{ if and .IsLoggin .IsAdmin }} admin {{ else if .IsLoggin }} login user {{ else }} guest{{end}}
`
	andFuncTpl, err := template.New("andtest").Parse(andTplStr)
	if err != nil {
		panic(err)
	}
	err = andFuncTpl.Execute(os.Stdout, mike)
	if err != nil {
		panic(err)
	}
	// the user is admin

	orTplStr := `
the user is{{ if or .IsLoggin .IsAdmin }} login {{ else }} guest{{end}}
`
	orFuncTpl, err := template.New("ortest").Parse(orTplStr)
	if err != nil {
		panic(err)
	}
	err = orFuncTpl.Execute(os.Stdout, mike)
	if err != nil {
		panic(err)
	}
	// the user is login

	notTplStr := `
the user is{{ if not .IsLoggin  }} guest{{ else }} login{{end}}
`
	notFuncTpl, err := template.New("nottest").Parse(notTplStr)
	if err != nil {
		panic(err)
	}
	err = notFuncTpl.Execute(os.Stdout, mike)
	if err != nil {
		panic(err)
	}
	// the user is login

	customFuncTplStr := `
{{ if hasPermission . "admin" }} the user is admin {{ end }}
`
	permissionFunc := func(user user, level string) bool {
		if level == "admin" {
			return user.IsAdmin
		}
		return false

	}
	customFuncTpl, err := template.New("customFuncTest").Funcs(template.FuncMap{
		"hasPermission": permissionFunc,
	}).Parse(customFuncTplStr)
	if err != nil {
		panic(err)
	}
	err = customFuncTpl.Execute(os.Stdout, mike)
	if err != nil {
		panic(err)
	}
	// the user is admin

}
