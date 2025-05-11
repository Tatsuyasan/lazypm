package pkgman

import "testing"

func TestNpmName(t *testing.T) {
	npm := NewNPM(".")
	if npm.Name() != "npm" {
		t.Errorf("Expected 'npm', got '%s'", npm.Name())
	}
}
