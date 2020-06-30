package pinkflamingo

import (
	"testing"
)

func TestFizzBuzz(t *testing.T) {
	tests := []struct {
		name           string
		n              int
		want           string
		isMultipleOf15 bool
	}{
		{
			name:           "when number is 25",
			n:              25,
			want:           "Buzz",
			isMultipleOf15: false,
		},
		{
			name:           "when number is 21",
			n:              21,
			want:           "Fizz",
			isMultipleOf15: false,
		},
		{
			name:           "when number is 15",
			n:              15,
			want:           "FizzBuzz",
			isMultipleOf15: true,
		},
		{
			name:           "when number is 75",
			n:              75,
			want:           "FizzBuzz",
			isMultipleOf15: true,
		},
		{
			name:           "when number is 100",
			n:              100,
			want:           "Buzz",
			isMultipleOf15: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := FizzBuzz(tt.n)
			if got != tt.want {
				t.Errorf("FizzBuzz() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.isMultipleOf15 {
				t.Errorf("FizzBuzz() got1 = %v, want %v", got1, tt.isMultipleOf15)
			}
		})
	}
}

func Test_isMultipleOf(t *testing.T) {
	type args struct {
		n       int
		divisor int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "when number is 100 and divisor is 5",
			args: args{
				n:       100,
				divisor: 5,
			},
			want: true,
		},
		{
			name: "when number is 20 and divisor is 3",
			args: args{
				n:       20,
				divisor: 3,
			},
			want: false,
		},
		{
			name: "when number is 42 and divisor is 3",
			args: args{
				n:       42,
				divisor: 3,
			},
			want: true,
		},
		{
			name: "when number is 63 and divisor is 5",
			args: args{
				n:       63,
				divisor: 5,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isMultipleOf(tt.args.n, tt.args.divisor); got != tt.want {
				t.Errorf("isMultipleOf() = %v, want %v", got, tt.want)
			}
		})
	}
}
