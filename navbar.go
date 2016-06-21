package main

const NAVBAR_TEMPLATE = `
	{{ define "navbar" }}
		<div id="navbar">
	    {{ range $k, $v := . }}
			{{if $v }}	
				<a class="navbar_item" href="{{ $v }}">{{ $k }}</a>
		    {{ end }}
		{{ end }}
		</div>
	{{ end }}
`

func buildNavbar(conf Config) map[string]string {
	return map[string]string{
		"Twitter":   conf.Links.Twitter,
		"Email":     conf.Links.Email,
		"Github":    conf.Links.Github,
		"Bitbucket": conf.Links.Bitbucket,
		"Gitlab":    conf.Links.Gitlab,
	}
}
