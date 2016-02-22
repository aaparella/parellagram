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

func buildStyles(stylesPath string) *Styles {
	styles, err := getDirContents(stylesPath)
	if err != nil {
		log.Panic("Error building styles : ", err)
	}
	return &Styles{
		StylePath: "./styles/",
		Files:     styles,
	}
}
