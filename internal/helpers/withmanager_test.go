package helpers

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/Tatsuyasan/lazyPm/internal/models"
)

type WithManagerTest struct {
	name          string
	pmFlag        string
	files         map[string]string
	expectedError bool
	expectedType  string
}

func TestWithManager(t *testing.T) {
	tests := []WithManagerTest{
		{
			name:   "With NPM",
			pmFlag: "npm",
			files: map[string]string{
				"package-lock.json": "",
			},
			expectedError: false,
			expectedType:  "*pkgman.NPM",
		},
		{
			name:   "With GoPM",
			pmFlag: "go",
			files: map[string]string{
				"go.mod": "",
			},
			expectedError: false,
			expectedType:  "*pkgman.GoPM",
		},
		{
			name:          "Unknown package manager",
			pmFlag:        "unknown",
			files:         map[string]string{},
			expectedError: true,
			expectedType:  "test",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a temporary directory
			tempDir, err := os.MkdirTemp("", "test-with-manager")
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

			// Mock callback function
			callback := func(pm models.PackageManager) error {
				if gotType := fmt.Sprintf("%T", pm); gotType != tt.expectedType {
					return fmt.Errorf("expected type %s, but got %s", tt.expectedType, gotType)
				}
				return nil
			}

			// Run the function
			err = WithManager(tt.pmFlag, callback)

			// Check for errors
			if tt.expectedError {
				if err == nil {
					t.Errorf("Expected an error but got none")
				}
				return
			}
			if err != nil {
				t.Errorf("Unexpected error: %v", err)
			}
		})
	}
}
