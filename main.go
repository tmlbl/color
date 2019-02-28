package main

import (
	"flag"
	"net/http"
	"text/template"
)

const pageTemplate = `
<head>
	<style>
		body {
			background-color: {{.color}};
			overflow: hidden;
		}

		.text {
			position: absolute;
			top: 30%;
			width: 100%;
		}

		h1 {
			color: {{.color}};
			filter: invert(100%);
			text-align: center;
			font-size: 4em;
			font-family: sans-serif;
		}
	</style>
</head>
<body>
	<div class="text">
		<h1>{{.text}}</h1>
	</div>
</body>
`

var (
	text  string
	color string
)

func main() {
	flag.StringVar(&text, "text", "Hello, world!", "Text to display")
	flag.StringVar(&color, "color", "cornflowerblue", "Color of the background")
	flag.Parse()

	tmpl, _ := template.New("page").Parse(pageTemplate)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl.Execute(w, map[string]string{
			"text":  text,
			"color": color,
		})
	})
	http.ListenAndServe(":80", nil)
}
