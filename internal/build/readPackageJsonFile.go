package build

import (
	"encoding/json"
	"os"
)

// PackageJSON represents the typical structure of a package.json file.
// This struct doesn't cover all possible fields but provides a starting point.
type PackageJSON struct {
	Name            string            `json:"name"`
	Version         string            `json:"version"`
	Description     string            `json:"description"`
	Main            string            `json:"main"`
	Scripts         map[string]string `json:"scripts"`
	Dependencies    map[string]string `json:"dependencies"`
	DevDependencies map[string]string `json:"devDependencies"`
}

// ReadPackageJSON reads a package.json file and returns its content as a Go struct.
func readPackageJSON(path string) (*PackageJSON, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var pkg PackageJSON
	err = json.Unmarshal(file, &pkg)
	if err != nil {
		return nil, err
	}

	return &pkg, nil
}
