package parser

import "testing"

func TestParsingWorker(t *testing.T) {
	tests := []struct {
		name  string
		input []byte
	}{
		{
			name:  "test 1",
			input: []byte{3, 0, 0, 0, 32, 4, 0, 1, 0, 10, 32, 55, 46, 50, 46, 49, 46, 50, 51, 73, 0, 2, 0, 8, 120, 114, 118, 57, 107, 45, 114, 49, 3, 0, 0, 0, 234, 3, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 192, 168, 80, 103, 0, 0, 19, 206, 57, 112, 1, 254, 94, 98, 129, 171, 0, 0, 215, 126, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 192, 168, 80, 128, 0, 179, 131, 152, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 0, 91, 1, 4, 19, 206, 0, 90, 192, 168, 8, 8, 62, 2, 6, 1, 4, 0, 1, 0, 1, 2, 6, 1, 4, 0, 1, 0, 4, 2, 6, 1, 4, 0, 1, 0, 128, 2, 2, 128, 0, 2, 2, 2, 0, 2, 6, 65, 4, 0, 0, 19, 206, 2, 20, 5, 18, 0, 1, 0, 1, 0, 2, 0, 1, 0, 2, 0, 2, 0, 1, 0, 128, 0, 2, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 0, 75, 1, 4, 19, 206, 0, 90, 57, 112, 1, 254, 46, 2, 44, 2, 0, 1, 4, 0, 1, 0, 1, 1, 4, 0, 2, 0, 1, 1, 4, 0, 1, 0, 4, 1, 4, 0, 2, 0, 4, 1, 4, 0, 1, 0, 128, 1, 4, 0, 2, 0, 128, 65, 4, 0, 0, 19, 206},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			parsingWorker(tt.input, nil)
		})
	}
}