package build

import (
	"github.com/tdewolff/minify/v2"
	"github.com/tdewolff/minify/v2/css"
	"log"
	"os"
	"path/filepath"
)

type Library struct {
	SourcePath       string
	DestinationPaths []string
}

func (receiver Library) Build() {
	// Read the original Build
	content, err := os.ReadFile(receiver.SourcePath)
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	// Create a new minifier
	m := minify.New()
	m.AddFunc("text/css", css.Minify)

	// Minify the Build content
	minifiedContent, err := m.Bytes("text/css", content)
	if err != nil {
		log.Fatalf("Error minifying Build: %v", err)
	}

	for _, path := range receiver.DestinationPaths {
		ensureDirectoryExists(path)
		// Write the minified content to the destination
		err = os.WriteFile(filepath.Join(path, getMinifiedFileName(receiver.SourcePath)), minifiedContent, 0644)
		// Write the unminified original content to the destination
		err = os.WriteFile(filepath.Join(path, filepath.Base(receiver.SourcePath)), content, 0644)
		if err != nil {
			log.Fatalf("Error writing minified CSS to file: %v", err)
		}
	}
}
