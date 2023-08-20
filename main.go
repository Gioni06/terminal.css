package main

import (
	"fmt"
	"github.com/Gioni06/terminalcss/internal/build"
	"github.com/fsnotify/fsnotify"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func main() {

	// Get the current working directory
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal("Error getting current working directory:", err)
	}

	// libPath is the path to the terminal.css library source
	libPath := filepath.Join(cwd, "lib", "terminal.css")

	// libPublishPath is the path to the terminal.css library publish directory that gets published to npm
	libPublishPath := filepath.Join(cwd, "dist")

	// docsPath is the path to the destination directory
	docsPath := filepath.Join(cwd, "public")

	// clear docsPath directory
	err = os.RemoveAll(docsPath)
	err = os.RemoveAll(libPublishPath)

	library := build.Library{
		SourcePath:       libPath,
		DestinationPaths: []string{libPublishPath, docsPath},
	}

	library.Build()

	layoutsPath := filepath.Join(cwd, "templates/layouts")
	partialsPath := filepath.Join(cwd, "templates/partials")
	pagesPath := filepath.Join(cwd, "pages")
	staticAssetsPath := filepath.Join(cwd, "static")

	docs := build.Docs{
		PackageJSONPath:  filepath.Join(cwd, "package.json"),
		DestinationPath:  docsPath,
		LayoutsPath:      layoutsPath,
		PartialsPath:     partialsPath,
		PagesPath:        pagesPath,
		StaticAssetsPath: staticAssetsPath,
	}

	err = docs.Build()
	if err != nil {
		panic(err)
	}

	// Check if program has the --serve flag
	if len(os.Args) == 2 && os.Args[1] == "--serve" {

		// Setup watchers for libPath and docsPath
		watcher, err := fsnotify.NewWatcher()
		if err != nil {
			log.Fatal(err)
		}
		defer watcher.Close()

		// Function to rebuild library and docs
		rebuild := func() {
			library.Build()
			err = docs.Build()
			if err != nil {
				log.Println("Error rebuilding docs:", err)
			}
		}

		// Watch for changes in libPath and docsPath
		pathsToWatch := []string{libPath, pagesPath, layoutsPath, partialsPath, staticAssetsPath}
		for _, path := range pathsToWatch {
			err = watcher.Add(path)
			if err != nil {
				log.Fatal(err)
			}
		}

		// Start a goroutine to handle events
		go func() {
			for {
				select {
				case event, ok := <-watcher.Events:
					if !ok {
						return
					}
					if event.Op&fsnotify.Write == fsnotify.Write {
						log.Println("Detected change, rebuilding...")
						rebuild()
					}
				case err, ok := <-watcher.Errors:
					if !ok {
						return
					}
					log.Println("Watcher error:", err)
				}
			}
		}()

		fmt.Println("Serving dist directory at http://localhost:8080. Press Ctrl+C to stop. Restart the program to rebuild the docs.")
		// serve dist directory at localhost:8080
		port := ":8080"
		handler := http.FileServer(http.Dir(docsPath))
		err = http.ListenAndServe(port, handler)
		if err != nil {
			return
		}
	}
}
