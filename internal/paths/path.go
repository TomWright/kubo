package paths

import (
	"path/filepath"
)

// StdPath allows us to use Paths directly via package func's.
var StdPath = NewPath()

// ConfigFile returns the path to the config file for the given service, environment and template.
func ConfigFile(service string, environment string, template string) string {
	return StdPath.ConfigFile(service, environment, template)
}

// TemplateDir returns the path to the given templates directory.
func TemplateDir(template string) string {
	return StdPath.TemplateDir(template)
}

// TemplateFile returns the path to the given template file.
func TemplateFile(template string, filename string) string {
	return StdPath.TemplateFile(template, filename)
}

// ManifestDir returns the path to the manifest directory of the given service and environment.
func ManifestDir(service string, environment string) string {
	return StdPath.ManifestDir(service, environment)
}

// ManifestFile returns the path to the given manifest file for the given service and environment.
func ManifestFile(service string, environment string, filename string) string {
	return StdPath.ManifestFile(service, environment, filename)
}

// NewPath returns a new instance of Paths.
func NewPath() *Paths {
	return new(Paths).SetBase(".")
}

// Paths contains a base path that we can use as a root for all other paths.
type Paths struct {
	base string
}

// SetBase sets the base path to be used when calculating other paths.
func (p *Paths) SetBase(path string) *Paths {
	p.base = filepath.Clean(path)
	return p
}

// ConfigFile returns the path to the config file for the given service, environment and template.
func (p Paths) ConfigFile(service string, environment string, template string) string {
	return filepath.Join(p.base, "config", service, environment, template+".yaml")
}

// TemplateDir returns the path to the given templates directory.
func (p Paths) TemplateDir(template string) string {
	return filepath.Join(p.base, "templates", template)
}

// TemplateFile returns the path to the given template file.
func (p Paths) TemplateFile(template string, filename string) string {
	return filepath.Join(p.TemplateDir(template), filename)
}

// ManifestDir returns the path to the manifest directory of the given service and environment.
func (p Paths) ManifestDir(service string, environment string) string {
	return filepath.Join(p.base, "k8s", service, environment)
}

// ManifestFile returns the path to the given manifest file for the given service and environment.
func (p Paths) ManifestFile(service string, environment string, filename string) string {
	return filepath.Join(p.ManifestDir(service, environment), filename)
}
