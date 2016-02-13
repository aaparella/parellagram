package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"text/template"
)

type DetailedPage struct {
	Style *Styles
	Post  *Post
}

const DETAILED_PAGE_TEMPLATE = `
		<html>
			<head>
				{{ template "styles" .Style }}
			</head>
			<body>
				<a href="/">
					<h1 id="header">Parellagram</h1>
				</a>
				{{ template "post" .Post }}
			</body>
		</html>
	`

func serveDetailedPage(w http.ResponseWriter, r *http.Request) {
	styles, _ := buildStyles()
	file, err := os.Open(path.Join("./posts", path.Base(r.URL.String())))
	if err != nil {
		log.Fatal(err)
	}
	post, err := createPost(file)
	if err != nil {
		log.Fatal(err)
	}

	detailed := DetailedPage{
		Style: styles,
		Post:  post,
	}
	buildDetailedPage(detailed, w)
}

func buildDetailedPage(page DetailedPage, w io.Writer) {
	tmpl := template.New("detailed")
	var err error

	parse := func(template string) {
		if err != nil {
			return
		}
		tmpl, err = tmpl.Parse(template)
	}

	parse(DETAILED_PAGE_TEMPLATE)
	parse(STYLES_TEMPLATE)
	parse(POST_TEMPLATE)

	if err != nil {
		log.Fatal(err)
	}
	_ = tmpl.Execute(w, page)
}
