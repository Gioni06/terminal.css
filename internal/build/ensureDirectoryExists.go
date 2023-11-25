package build

import (
	"log"
	"os"
)

// ensureDirectoryExists ensures that the directory at the given path exists and creates it if it doesn't.
func ensureDirectoryExists(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.MkdirAll(path, 0755)
		if err != nil {
			log.Fatalf("Error creating directory: %v", err)
		}
	}
}
