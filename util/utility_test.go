package util

import (
	"reflect"
	"testing"
)

func TestIsEquals(t *testing.T) {
	tests := []struct {
		input     interface{}
		testInput interface{}
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
		},
		{
			input:     1,
			testInput: 1,
			expected:  true,
		},
		{
			input:     1,
			testInput: 12,
			expected:  false,
		},
	}

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
		name  string
		input string
		want  int
	}{
		{
			name:  "when input is 1",
			input: "1",
			want:  1,
		},
		{
			name:  "when input is 3",
			input: "3",
			want:  3,
		},
		{
			name:  "when input is 4ff",
			input: "4ff",
			want:  0,
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
		name  string
		input string
		want  bool
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
			if got := IsValidData(tt.input); got != tt.want {
				t.Errorf("IsValidData() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsSmaller(t *testing.T) {
	type args struct {
		firstNumber  int
		secondNumber int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "when firstNumber is 100 and secondNumber is 5",
			args: args{
				firstNumber:  100,
				secondNumber: 5,
			},
			want: false,
		},
		{
			name: "when firstNumber is 100 and secondNumber is 200",
			args: args{
				firstNumber:  100,
				secondNumber: 200,
			},
			want: true,
		},
		{
			name: "when firstNumber is 100 and secondNumber is 100",
			args: args{
				firstNumber:  100,
				secondNumber: 100,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsSmaller(tt.args.firstNumber, tt.args.secondNumber); got != tt.want {
				t.Errorf("IsSmaller() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsValidData(t *testing.T) {
	tests := []struct {
		name string
		data string
		want bool
	}{
		{
			name: "when data is emptyString",
			data: "",
			want: false,
		},
		{
			name: "when data contains value",
			data: "ffd",
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsValidData(tt.data); got != tt.want {
				t.Errorf("IsValidData() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSplitBySeparator(t *testing.T) {
	type args struct {
		data      string
		separator string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "when data is XY * Y and separator is space",
			args: args{
				data:  "XY * Y",
				separator: " ",
			},
			want: []string{"XY", "*", "Y"},
		},
		{
			name: "when data is XY* Y and separator is space",
			args: args{
				data:  "XY* Y",
				separator: " ",
			},
			want: []string{"XY*", "Y"},
		},
		{
			name: "when data is XY * Y and separator is comma",
			args: args{
				data:  "XY * Y",
				separator: ",",
			},
			want: []string{"XY * Y"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SplitBySeparator(tt.args.data, tt.args.separator); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SplitBySeparator() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToString(t *testing.T) {
	tests := []struct {
		name string
		data int
		want string
	}{
		{
			name: "when data is 1",
			data: 1,
			want: "1",
		},
		{
			name: "when data is 100",
			data: 100,
			want: "100",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToString(tt.data); got != tt.want {
				t.Errorf("ToString() = %v, want %v", got, tt.want)
			}
		})
	}
}
