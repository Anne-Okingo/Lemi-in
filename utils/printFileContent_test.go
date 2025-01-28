package utils

import (
	"fmt"
	"os"
	"reflect"
	"testing"
)

func TestValidContent(t *testing.T) {
	// create temporary test case

	tests := []struct {
		name          string
		content       string
		expected      []string
		expectederror bool
	}{
		{
			name: "with empty lines",
			content: `20
		##start
		0 1 2

		1 2 3
		2 5 6

		##end
		3 4 6
		`,
			expected:      []string{"20", "##start", "0 1 2", "1 2 3", "2 5 6", "##end", "3 4 6"},
			expectederror: false,
		},
		{
			name: "with comments",
			content: `20
	##start
	0 1 2
	#comment
	1 2 3
	2 5 6
	#another comment
	##end
	3 4 6
	`,
			expected:      []string{"20", "##start", "0 1 2", "1 2 3", "2 5 6", "##end", "3 4 6"},
			expectederror: false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			tempFile, err := os.CreateTemp("", "validcontentfile")
			if err != nil {
				t.Fatal("failed to create temp file", err)
			}
			defer os.Remove(tempFile.Name())

			if _, err := tempFile.WriteString(test.content); err != nil {
				t.Fatal("Error writting content to temp file", err)
			}
			tempFile.Close()

			finalContent, err := ValidContent(tempFile.Name())
			if (err != nil) || !reflect.DeepEqual(test.expected, finalContent) {
				t.Fatal(fmt.Printf("Error Test ValidContent failed expected%v, got%v", test.expected, finalContent))
			}
		})
	}
}
