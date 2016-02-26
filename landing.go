package main

import (
	"io"
	"log"
	"os"
	"path"
	"text/template"
)

type LandingPage struct {
	Title  string
	Styles *Styles
	Posts  []*Post
}

const LANDING_PAGE_TEMPLATE = `
	<html>
		<title>{{ .Title }}</title>
		<head>
			{{ template "styles" .Styles }}
		</head>
		<body>
			<h1 id="header">Parellagram</h1>
			{{ template "posts-preview" .Posts }}
		</div>
		</body>
	</html>
	`

func buildLandingPage(page LandingPage, w io.Writer) {
	tmpl := template.New("page")
	var err error

	parse := func(template string) {
		if err != nil {
			return
		}
		tmpl, err = tmpl.Parse(template)
	}

	parse(LANDING_PAGE_TEMPLATE)
	parse(STYLES_TEMPLATE)
	parse(POSTS_PREVIEW_TEMPLATE)
	parse(POST_PREVIEW_TEMPLATE)

	if err != nil {
		log.Fatal(err)
	}

	_ = tmpl.Execute(w, page)
}

func saveLandingPage(conf Config) {
	posts := buildPosts(conf.Resources.Posts)
	styles := buildStyles(conf.Resources.Styles)
	page := LandingPage{
		Title:  conf.Website.Title,
		Styles: styles,
		Posts:  posts,
	}
	p := path.Join(os.TempDir(), "parellagram", "index.html")
	file, err := os.Create(p)
	if err != nil {
		log.Fatal("Error creating landing page : ", err)
	}
	buildLandingPage(page, file)
}
