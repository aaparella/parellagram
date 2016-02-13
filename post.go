package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"path"
	"regexp"
)

type Post struct {
	Title    string
	Date     string // TODO: Use Time struct
	Preview  string
	Content  string
	Filename string
}

const POST_PREVIEW_TEMPLATE = `
	{{ define "post-preview" }}
		<div id="post_preview">
		<a href="/post/{{ .Filename }}">
			<div class="post_title">{{ .Title }}</div>
			<div class="post_date">{{ .Date }}</div>
		</a>
		</div>
	{{ end }}
	`

const POSTS_PREVIEW_TEMPLATE = `
	{{ define "posts-preview" }}
		<div id="posts">
		{{ range . }}
			{{ template "post-preview" . }}
		{{ end }}
		</div>
	{{ end }}
	`

const POST_TEMPLATE = `
	{{ define "post" }}
		<div id="post">
			<div id="post_title">{{ .Title }}</div>
			<div id="post_date">{{ .Date }}</div>
			<br /><br />
			<div id="post_body">{{ .Content }}</div>
		</div>
	{{ end }}
	`

func createPost(file *os.File) (*Post, error) {
	post := &Post{}
	var err error = nil
	reader := bufio.NewReader(file)

	parse := func(prefix string) string {
		regex := regexp.MustCompile(prefix + ":")
		line, _ := reader.ReadString('\n')
		if match := regex.FindStringIndex(line); match != nil {
			return line[match[1] : len(line)-1]
		}
		err = fmt.Errorf("Expected %s string, got : \"%s\"", prefix, line)
		return ""
	}

	post.Title = parse("Title")
	post.Date = parse("Date")
	post.Filename = path.Base(file.Name())

	buffer := bytes.NewBuffer(nil)
	io.Copy(buffer, reader)
	post.Content = string(buffer.Bytes())

	return post, err
}

func buildPosts() ([]*Post, error) {
	filenames, err := getDirContents("posts")
	if err != nil {
		return nil, err
	}
	posts := make([]*Post, 0)
	for _, filename := range filenames {
		file, err := os.Open(path.Join("posts", filename))
		if err != nil {
			log.Println("Problem opening file : ", filename, ", ", err, " skipping")
		}
		post, err := createPost(file)
		if err != nil {
			log.Println("Problem parsing file : ", filename, ", ", err, " skipping")
		} else {
			posts = append(posts, post)
		}
	}
	return posts, nil
}
