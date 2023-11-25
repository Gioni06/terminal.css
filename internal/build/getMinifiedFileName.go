package build

import (
	"fmt"
	"path/filepath"
	"strings"
)

func getMinifiedFileName(path string) string {
	// Extract the filename without the preceding directories.
	fileName := filepath.Base(path)

	// Split the filename into name and extension.
	ext := filepath.Ext(fileName)
	name := strings.TrimSuffix(fileName, ext)

	// If the filename already ends with ".min", don't add another ".min" to it.
	if strings.HasSuffix(name, ".min") {
		return fileName
	}

	// Append ".min" to the name and concatenate with the extension.
	minifiedFileName := fmt.Sprintf("%s.min%s", name, ext)

	return minifiedFileName
}
