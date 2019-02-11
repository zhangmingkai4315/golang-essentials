package main

import (
	"flag"
	"html/template"
	"log"
	"net/http"
)

var addr = flag.String("addr", ":1718", "http web sever port")

const templateStr = `
<html>
<head>
<title>QR Link Generator</title>
</head>
<body>
{{if .}}
<img src="http://chart.apis.google.com/chart?chs=300x300&cht=qr&choe=UTF-8&chl={{.}}" />
<br>
{{.}}
<br>
<br>
{{end}}
<form action="/" name=f method="GET"><input maxLength=1024 size=70
name=s value="" title="Text to QR Encode"><input type=submit
value="Show QR" name=qr>
</form>
</body>
</html>
`

var templ = template.Must(template.New("qr").Parse(templateStr))

func qrhandler(w http.ResponseWriter, req *http.Request) {
	templ.Execute(w, req.FormValue("s"))
}

func main() {
	flag.Parse()
	http.Handle("/", http.HandlerFunc(qrhandler))
	log.Printf("start server at %s", *addr)
	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatalf("start server error : %f", err)
	}

}
