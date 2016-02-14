package main

import (
	"io"
	"log"
	"net/http"
	"text/template"
)

type LandingPage struct {
	Styles *Styles
	Posts  []*Post
}

const LANDING_PAGE_TEMPLATE = `
	<html>
		<title>Parellagram</title>
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

func serveLandingPage(w http.ResponseWriter, r *http.Request) {
	posts := buildPosts()
	styles := buildStyles()
	page := LandingPage{
		Styles: styles,
		Posts:  posts,
	}
	buildLandingPage(page, w)
}
