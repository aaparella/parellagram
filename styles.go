package main

import "log"

type Styles struct {
	StylePath string
	Files     []string
}

const STYLES_TEMPLATE = `
	{{ define "styles" }}
		{{ range .Files }}
			<link rel="stylesheet" href="/styles/{{.}}" type="text/css">
		{{ end }}
	{{ end }}
	`

func buildStyles() *Styles {
	styles, err := getDirContents("styles")
	if err != nil {
		log.Panic("Error building styles : ", err)
	}
	return &Styles{
		StylePath: "./styles/",
		Files:     styles,
	}
}
