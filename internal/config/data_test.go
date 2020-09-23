package config_test

import (
	"github.com/google/go-cmp/cmp"
	"github.com/tomwright/kubo/internal/config"
	"testing"
)

const testStringA = `
name: test
environment: development
image:
  name: tomwright/test
  version: v1.0.0
http:
  port: 8000
grpc:
  port: 9000
env:
  - name: A
    value: auth-cluster-ip
  - name: B
    value: asdasd
`

var testDataA = config.Data{
	"name":        "test",
	"environment": "development",
	"image": config.Data{
		"name":    "tomwright/test",
		"version": "v1.0.0",
	},
	"http": config.Data{
		"port": 8000,
	},
	"grpc": config.Data{
		"port": 9000,
	},
	"env": []interface{}{
		config.Data{"name": "A", "value": "auth-cluster-ip"},
		config.Data{"name": "B", "value": "asdasd"},
	},
}

func TestData_Set(t *testing.T) {
	a := &config.Data{
		"name":        "test",
		"environment": "development",
		"image": config.Data{
			"name":    "tomwright/test",
			"version": "v1.0.0",
		},
		"http": config.Data{
			"port": 8000,
		},
		"grpc": config.Data{
			"port": 9000,
		},
	}

	if err := a.Set("name", "tester"); err != nil {
		t.Errorf("unexpected error: %s", err)
	}
	if err := a.Set("http.port", 8080); err != nil {
		t.Errorf("unexpected error: %s", err)
	}
	if err := a.Set("some.other.group", "here"); err != nil {
		t.Errorf("unexpected error: %s", err)
	}

	if exp, got := "tester", a.StringR("name"); exp != got {
		t.Errorf("expected name %v, got %v", exp, got)
	}
	if exp, got := 8080, a.DataR("http").IntR("port"); exp != got {
		t.Errorf("expected http port %v, got %v", exp, got)
	}
	if exp, got := "here", a.DataR("some").DataR("other").StringR("group"); exp != got {
		t.Errorf("expected some other group %v, got %v", exp, got)
	}
}

func TestData_Values(t *testing.T) {
	if exp, got := "test", testDataA.StringR("name"); exp != got {
		t.Errorf("expected name %v, got %v", exp, got)
	}
	if exp, got := "development", testDataA.StringR("environment"); exp != got {
		t.Errorf("expected environment %v, got %v", exp, got)
	}
	if exp, got := "tomwright/test", testDataA.DataR("image").StringR("name"); exp != got {
		t.Errorf("expected image name %v, got %v", exp, got)
	}
	if exp, got := "v1.0.0", testDataA.DataR("image").StringR("version"); exp != got {
		t.Errorf("expected image version %v, got %v", exp, got)
	}
	if exp, got := 8000, testDataA.DataR("http").IntR("port"); exp != got {
		t.Errorf("expected http port %v, got %v", exp, got)
	}
	if exp, got := 9000, testDataA.DataR("grpc").IntR("port"); exp != got {
		t.Errorf("expected grpc port %v, got %v", exp, got)
	}
}

func TestFromBytes(t *testing.T) {
	d, err := config.FromBytes([]byte(testStringA))
	if err != nil {
		t.Errorf("unexpected error: %s", err)
		return
	}

	if !cmp.Equal(d, testDataA) {
		t.Errorf("unexpected data:\n%s\n", cmp.Diff(d, testDataA))
	}
}
