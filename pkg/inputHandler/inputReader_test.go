package inputHandler

import (
	"bufio"
	"errors"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetInput(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		expected    string
		shouldError bool
		mockedError error
	}{
		{
			name:        "Valid input",
			input:       "Hello, World!\n",
			expected:    "Hello, World!",
			shouldError: false,
		},
		{
			name:        "Input with extra spaces",
			input:       "   GoLang   \n",
			expected:    "GoLang",
			shouldError: false,
		},
		{
			name:        "Empty input",
			input:       "\n",
			expected:    "",
			shouldError: false,
		},
		{
			name:        "Error during reading input",
			input:       "",
			expected:    "",
			shouldError: true,
			mockedError: errors.New("mocked error"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var reader *bufio.Reader
			if test.mockedError != nil {
				// Simulating an error case (error override)
				reader = bufio.NewReader(&errorReader{err: test.mockedError})
			} else {
				// Simulating valid input case
				reader = bufio.NewReader(strings.NewReader(test.input))
			}

			ir := NewInputReader(reader)

			result, err := ir.GetInput("Enter input: ")
			if test.shouldError {
				assert.Error(t, err)
				assert.Equal(t, test.mockedError.Error(), err.Error())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, test.expected, result)
			}
		})
	}
}

// Helper: Mocking an io.Reader that returns an error
type errorReader struct {
	err error
}

func (e *errorReader) Read(p []byte) (n int, err error) {
	return 0, e.err
}
