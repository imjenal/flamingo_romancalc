package util

import (
	"testing"
)

func TestIsEquals(t *testing.T) {
	tests := []struct {
		input     string
		testInput string
		expected  bool
	}{
		{
			input:     "Hello",
			testInput: "Hello",
			expected:  true,
		},
		{
			input:     "Hello",
			testInput: "hello",
			expected:  false,
		}}

	for testNumber, test := range tests {
		actual := IsEquals(test.input, test.testInput)
		if actual != test.expected {
			t.Errorf("Test No %d: Expected: %v but Got %v", testNumber+1, test.expected, actual)
		}
	}
}

func TestIsEmpty(t *testing.T) {
	tests := []struct {
		data     string
		expected bool
	}{
		{
			data:     "",
			expected: true,
		},
		{
			data:     "Hello",
			expected: false,
		},
	}

	for testNumber, test := range tests {
		actual := IsEmpty(test.data)
		if actual != test.expected {
			t.Errorf("Test No %d: Expected: %v but Got %v", testNumber+1, test.expected, actual)
		}
	}
}

func TestToInteger(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    int
	}{
		{
			name:    "when input is 1",
			input:   "1",
			want:    1,
		},
		{
			name:    "when input is 3",
			input:   "3",
			want:    3,
		},
		{
			name:    "when input is 4ff",
			input:   "4ff",
			want:    0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := ToInteger(tt.input)
			if got != tt.want {
				t.Errorf("ToInteger() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToLowerCase(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "when input is Hello",
			input: "Hello",
			want:  "hello",
		},
		{
			name:  "when input is Hola",
			input: "Hola",
			want:  "hola",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToLowerCase(tt.input); got != tt.want {
				t.Errorf("ToLowerCase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsValidNumber(t *testing.T) {

	tests := []struct {
		name string
		input string
		want bool
	}{
		{
			name:  "when input is empty",
			input: "",
			want:  false,
		},
		{
			name:  "when input is not empty",
			input: "33",
			want:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsValidNumber(tt.input); got != tt.want {
				t.Errorf("IsValidNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}