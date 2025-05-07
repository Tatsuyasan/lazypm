package helpers

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

type AutodetectTest struct {
	name          string
	files         map[string]string
	expectedError bool
	expectedType  string
}

func TestDetectPackageManager(t *testing.T) {
	tests := []AutodetectTest{
		{
			name: "Detect NPM",
			files: map[string]string{
				"package-lock.json": "",
			},
			expectedError: false,
			expectedType:  "*pkgman.NPM",
		},
		{
			name: "Detect GoPM",
			files: map[string]string{
				"go.mod": "",
			},
			expectedError: false,
			expectedType:  "*pkgman.GoPM",
		},
		{
			name:          "No package manager",
			files:         map[string]string{},
			expectedError: true,
			expectedType:  "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a temporary directory
			tempDir, err := os.MkdirTemp("", "test-detect-pm")
			if err != nil {
				t.Fatalf("Failed to create temp directory: %v", err)
			}
			defer os.RemoveAll(tempDir)

			// Create test files
			for name, content := range tt.files {
				filePath := filepath.Join(tempDir, name)
				if err := os.WriteFile(filePath, []byte(content), 0644); err != nil {
					t.Fatalf("Failed to create test file %s: %v", name, err)
				}
			}

			// Run the function
			pm, err := DetectPackageManager(tempDir)

			// Check for errors
			if tt.expectedError {
				if err == nil {
					t.Errorf("Expected an error but got none")
				}
				return
			}
			if err != nil {
				t.Errorf("Unexpected error: %v", err)
				return
			}

			// Check the type of the returned package manager
			if gotType := fmt.Sprintf("%T", pm); gotType != tt.expectedType {
				t.Errorf("Expected type %s, but got %s", tt.expectedType, gotType)
			}
		})
	}
}
