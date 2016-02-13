package main

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

func buildStyles() (*Styles, error) {
	styles, err := getDirContents("styles")
	if err != nil {
		return nil, err
	}
	return &Styles{
		StylePath: "./styles/",
		Files:     styles,
	}, nil
}
