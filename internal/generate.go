package internal

import (
	"fmt"
	"github.com/tomwright/kubo/internal/config"
	"github.com/tomwright/kubo/internal/paths"
	"io"
	"io/ioutil"
	"os"
	"text/template"
)

// GenerateTemplateFile prepares and generates the given file.
func GenerateTemplateFile(f os.FileInfo, serviceName string, environment string, serviceType string, config config.Data) error {
	if f.IsDir() {
		return nil
	}

	inputFile, err := os.Open(paths.TemplateFile(serviceType, f))
	if err != nil {
		return fmt.Errorf("could not open input file: %w", err)
	}
	defer inputFile.Close()

	outputFile, err := os.Create(paths.ManifestFile(serviceName, environment, f))
	if err != nil {
		return fmt.Errorf("could not create output file: %w", err)
	}
	defer outputFile.Close()

	if err := Generate(inputFile, outputFile, config); err != nil {
		return fmt.Errorf("could not generate template file: %w", err)
	}

	return nil
}

// Generate takes a template as input, executes it with the config as the values
// and writes to the output.
func Generate(input io.Reader, output io.Writer, config config.Data) error {
	t := template.New("generate")
	inputBytes, err := ioutil.ReadAll(input)
	if err != nil {
		return fmt.Errorf("could not read input: %w", err)
	}

	t, err = t.Parse(string(inputBytes))
	if err != nil {
		return fmt.Errorf("could not parse input template: %w", err)
	}

	if err := t.Execute(output, config); err != nil {
		return fmt.Errorf("could not execute template: %w", err)
	}

	return nil
}
