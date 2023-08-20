package build

import (
	"bytes"
	"gopkg.in/yaml.v2"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
)

type Docs struct {
	PackageJSONPath  string
	DestinationPath  string
	LayoutsPath      string
	PartialsPath     string
	PagesPath        string
	StaticAssetsPath string
}

type FrontMatter struct {
	Layout      string `yaml:"layout"`
	Title       string `yaml:"title"`
	Description string `yaml:"description"`
}

func (receiver Docs) Build() error {
	// Read Markdown files from PagesPath
	files, err := os.ReadDir(receiver.PagesPath)
	if err != nil {
		return err
	}

	pgkJSON, _ := readPackageJSON(receiver.PackageJSONPath)
	version := pgkJSON.Version

	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".md") {
			content, err := os.ReadFile(filepath.Join(receiver.PagesPath, file.Name()))
			if err != nil {
				return err
			}

			// Parse front matter
			frontMatter, markdownContent := parseFrontMatter(content)

			// Load and parse the layout template with partials
			tmpl, err := parseTemplateWithPartials(receiver.LayoutsPath, receiver.PartialsPath, frontMatter.Layout)
			if err != nil {
				return err
			}

			// Convert Markdown to HTML
			renderer := html.NewRenderer(html.RendererOptions{})
			htmlContent := markdown.ToHTML([]byte(markdownContent), nil, renderer)

			// Execute the template with the HTML content
			var buf bytes.Buffer
			err = tmpl.Execute(&buf, map[string]interface{}{
				"Content":     string(htmlContent),
				"FrontMatter": frontMatter,
				"Version":     version,
			})
			if err != nil {
				return err
			}

			if file.Name() == "index.md" {
				err = os.WriteFile(filepath.Join(receiver.DestinationPath, "index.html"), buf.Bytes(), 0644)
				if err != nil {
					return err
				}
				continue
			} else {
				// Create directory for the file
				err = os.MkdirAll(filepath.Join(receiver.DestinationPath, strings.Replace(file.Name(), ".md", "", -1)), 0755)
				if err != nil {
					return err
				}
				// Write to the destination path
				err = os.WriteFile(filepath.Join(receiver.DestinationPath, strings.Replace(file.Name(), ".md", "/index.html", -1)), buf.Bytes(), 0644)
				if err != nil {
					return err
				}
			}
		}
	}

	// Copy static assets
	err = copyDir(receiver.StaticAssetsPath, filepath.Join(receiver.DestinationPath))

	return nil
}

func copyDir(sourcePath string, destinationPath string) error {

	ensureDirectoryExists(destinationPath)
	// Read source directory
	files, err := os.ReadDir(sourcePath)
	if err != nil {
		return err
	}

	for _, file := range files {
		if file.IsDir() {
			// Recursively copy subdirectories
			err = copyDir(filepath.Join(sourcePath, file.Name()), filepath.Join(destinationPath, file.Name()))
			if err != nil {
				return err
			}
		} else {
			// Copy file
			readFile, err := os.ReadFile(filepath.Join(sourcePath, file.Name()))
			if err != nil {
				return err
			}
			err = os.WriteFile(filepath.Join(destinationPath, file.Name()), readFile, 0644)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func parseFrontMatter(content []byte) (FrontMatter, string) {
	// Split the content by the end of the front matter (assuming it's delineated by '---')
	parts := strings.SplitN(string(content), "---", 3)

	var frontMatter FrontMatter
	if len(parts) > 2 {
		err := yaml.Unmarshal([]byte(parts[1]), &frontMatter)
		if err != nil {
			return FrontMatter{}, ""
		}
		return frontMatter, parts[2]
	}

	return frontMatter, string(content)
}

func parseTemplateWithPartials(layoutsPath, partialsPath, layoutName string) (*template.Template, error) {
	// Parse main layout
	mainLayoutPath := filepath.Join(layoutsPath, layoutName+".html")
	tmpl, err := template.ParseFiles(mainLayoutPath)
	if err != nil {
		return nil, err
	}

	// Parse partials
	partials, err := os.ReadDir(partialsPath)
	if err != nil {
		return nil, err
	}

	for _, partial := range partials {
		_, err := tmpl.ParseFiles(filepath.Join(partialsPath, partial.Name()))
		if err != nil {
			return nil, err
		}
	}

	return tmpl, nil
}
