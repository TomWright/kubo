package internal_test

import (
	"github.com/google/go-cmp/cmp"
	"github.com/tomwright/kubo/internal"
	"testing"
)

func TestOverrideFlag_Set(t *testing.T) {
	overrides := make(internal.OverrideFlag, 0)
	if err := overrides.Set("a=asd"); err != nil {
		t.Errorf("unexpected error: %s", err)
	}
	if err := overrides.Set("b="); err != nil {
		t.Errorf("unexpected error: %s", err)
	}
	if err := overrides.Set("c=a=b"); err != nil {
		t.Errorf("unexpected error: %s", err)
	}
	if err := overrides.Set("a.b.c=d.e.f"); err != nil {
		t.Errorf("unexpected error: %s", err)
	}
	if err := overrides.Set(""); err != nil {
		t.Errorf("unexpected error: %s", err)
	}

	exp := internal.OverrideFlag{
		{
			Path:  "a",
			Value: "asd",
		},
		{
			Path:  "b",
			Value: "",
		},
		{
			Path:  "c",
			Value: "a=b",
		},
		{
			Path:  "a.b.c",
			Value: "d.e.f",
		},
	}

	if !cmp.Equal(exp, overrides) {
		t.Errorf("unexpected data:\n%s\n", cmp.Diff(exp, overrides))
	}
}
