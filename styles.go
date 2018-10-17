package main

import "log"

type Styles struct {
	StylePath string
	Files     []string
}

const STYLES_TEMPLATE = `
	{{ define "styles" }}
	<style>
		html {
			font-family: sans-serif;
		}

		#header {
			text-align: center;
			padding: 20px, 20px, 10px, 20px;
		}

		a #header:hover
		{
			color: white;
			text-shadow: 0 0 10px #000;
			-moz-text-shadow: 0 0 2px #000;
			-webkit-text-shadow: 0 0 2px #000;
		}

		#navbar {
			text-align: center;
			display: block;
		}

		.navbar_item {
			font-size: 20px;
			padding: 10px;
		}

		.navbar_item:hover {
			text-decoration: underline;
		}
		#post_preview {
			text-decoration: none;
			text-align: center;
			padding: 20px;

			margin-left: auto;
			margin-right: auto;

			margin-top: 10px;
			margin-bottom: 10px;

			border-style: solid;
			display: block;
			border-radius: 10px;
			width: 50%;
		}

		#post_preview .post_title {
			font-size: 24px;
		}

		#post_preview .post_date {
			text-decoration: italic;
		}

		a {
			text-decoration: none;
			color: black;
		}

		#post_preview:hover {
			background-color: #ffff99;
		}

		#post {
			margin-left: auto;
			margin-right: auto;
			width: 60%;
			font-family: sans-serif;
		}

		#post_title {
			font-size: 24px;
			text-align: center;
		}

		#post_date {
			text-align: center;
		}


		#post_body {
			text-align: left;
		}

		#post_body a {
			text-decoration: underline;
		}
	</style>
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
