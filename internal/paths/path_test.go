package paths_test

import (
	"github.com/tomwright/kubo/internal/paths"
	"testing"
)

var customBase = paths.NewPath().SetBase("/var/kubo")

func TestConfigFile(t *testing.T) {
	if exp, got := "config/ser/env/tem.yaml", paths.ConfigFile("ser", "env", "tem"); exp != got {
		t.Errorf("expected path of `%v`, got `%v`", exp, got)
	}
	if exp, got := "/var/kubo/config/ser/env/tem.yaml", customBase.ConfigFile("ser", "env", "tem"); exp != got {
		t.Errorf("expected path of `%v`, got `%v`", exp, got)
	}
}

func TestTemplateDir(t *testing.T) {
	if exp, got := "templates/tem", paths.TemplateDir("tem"); exp != got {
		t.Errorf("expected path of `%v`, got `%v`", exp, got)
	}
	if exp, got := "/var/kubo/templates/tem", customBase.TemplateDir("tem"); exp != got {
		t.Errorf("expected path of `%v`, got `%v`", exp, got)
	}
}

func TestTemplateFile(t *testing.T) {
	if exp, got := "templates/tem/a.yaml", paths.TemplateFile("tem", "a.yaml"); exp != got {
		t.Errorf("expected path of `%v`, got `%v`", exp, got)
	}
	if exp, got := "/var/kubo/templates/tem/a.yaml", customBase.TemplateFile("tem", "a.yaml"); exp != got {
		t.Errorf("expected path of `%v`, got `%v`", exp, got)
	}
}

func TestManifestDir(t *testing.T) {
	if exp, got := "k8s/ser/env", paths.ManifestDir("ser", "env"); exp != got {
		t.Errorf("expected path of `%v`, got `%v`", exp, got)
	}
	if exp, got := "/var/kubo/k8s/ser/env", customBase.ManifestDir("ser", "env"); exp != got {
		t.Errorf("expected path of `%v`, got `%v`", exp, got)
	}
}

func TestManifestFile(t *testing.T) {
	if exp, got := "k8s/ser/env/a.yaml", paths.ManifestFile("ser", "env", "a.yaml"); exp != got {
		t.Errorf("expected path of `%v`, got `%v`", exp, got)
	}
	if exp, got := "/var/kubo/k8s/ser/env/a.yaml", customBase.ManifestFile("ser", "env", "a.yaml"); exp != got {
		t.Errorf("expected path of `%v`, got `%v`", exp, got)
	}
}
