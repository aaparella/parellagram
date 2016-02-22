package main

import (
	"io"
	"log"
	"os"
	"path"
	"text/template"
)

type DetailedPage struct {
	Title string
	Style *Styles
	Post  *Post
}

const DETAILED_PAGE_TEMPLATE = `
		<html>
			<head>
				{{ template "styles" .Style }}
				<title>{{ .Title }} - {{ .Post.Title }}</title>
			</head>
			<body>
				<a href="/">
					<h1 id="header">{{ .Title }}</h1>
				</a>
				{{ template "post" .Post }}
			</body>
		</html>
	`

func saveDetailedPage(page DetailedPage, conf Config) {
	p := path.Join(conf.Artifacts.Path, conf.Resources.Posts, page.Post.Filename)
	file, err := os.Create(p)
	if err != nil {
		log.Fatal(err)
	}
	buildDetailedPage(page, file)
}

func saveDetailedPages(conf Config) {
	posts := buildPosts(conf.Resources.Posts)
	styles := buildStyles(conf.Resources.Styles)
	for _, post := range posts {
		page := DetailedPage{
			Title: conf.Website.Title,
			Style: styles,
			Post:  post,
		}

		saveDetailedPage(page, conf)
	}
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
