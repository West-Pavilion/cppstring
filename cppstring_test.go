package cppstring

import "testing"

func TestReverse(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Empty string",
			input:    "",
			expected: "",
		},
		{
			name:     "Single character",
			input:    "a",
			expected: "a",
		},
		{
			name:     "Basic English string",
			input:    "hello",
			expected: "olleh",
		},
		{
			name:     "String with spaces",
			input:    "hello world",
			expected: "dlrow olleh",
		},
		{
			name:     "Chinese characters",
			input:    "你好",
			expected: "好你",
		},
		{
			name:     "Mixed characters",
			input:    "Hello世界123",
			expected: "321界世olleH",
		},
		{
			name:     "String with emojis",
			input:    "👋🌍",
			expected: "🌍👋",
		},
		{
			name:     "Palindrome",
			input:    "racecar",
			expected: "racecar",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Reverse(tt.input)
			if result != tt.expected {
				t.Errorf("Reverse(%q) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}

func BenchmarkReverse(b *testing.B) {
	testString := "Hello, 世界! This is a test string with mixed characters 测试字符串 🌍👋"
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Reverse(testString)
	}
}