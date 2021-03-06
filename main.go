package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"path"
	"strings"
	"syscall"

	fswatch "github.com/andreaskoch/go-fswatch"
)

func skipDotFilesAndFolders(name string) bool {
	return strings.HasPrefix(path.Base(name), ".")
}

func watch(conf Config) {
	watcher := fswatch.NewFolderWatcher(conf.Resources.Posts, true, skipDotFilesAndFolders, 10)
	watcher.Start()

	for watcher.IsRunning() {
		modified := <-watcher.Modified()
		if modified {
			log.Println("Change detected, rebuilding pages...")
			buildWebsitePages(conf)
			log.Println("Done rebuilding pages")
		}
	}
}

func buildWebsitePages(conf Config) {
	clearTempDirectory(conf)
	saveDetailedPages(conf)
	saveLandingPage(conf)
}

func main() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT)

	go func() {
		_ = <-sigs
		log.Println("Cleaning up...")
		deleteTempDirectory()
		os.Exit(0)
	}()

	conf := getConfig()
	buildWebsitePages(conf)

	go watch(conf)

	http.Handle("/", http.FileServer(http.Dir(path.Join(os.TempDir(), "parellagram"))))
	http.HandleFunc("/uptime", Uptime)
	http.Handle("/styles/", http.StripPrefix("/styles/", http.FileServer(http.Dir(conf.Resources.Styles))))
	http.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir(conf.Resources.Images))))

	log.Println("Listening on port :", conf.Website.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", conf.Website.Port), nil))
}
