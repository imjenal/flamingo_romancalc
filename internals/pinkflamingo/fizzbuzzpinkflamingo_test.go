package pinkflamingo

import (
	"reflect"
	"testing"
)

func Test_convertToIntegers(t *testing.T) {
	type args struct {
		to     string
		from   string
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 int
	}{
		{
			name: "when to is 2 and from is 1",
			args: args{
				to: "2",
				from:"1",
			},
			want: 1,
			want1: 2,

		},
		{
			name: "when to is 7 and from is 1",
			args: args{
				to: "7",
				from:"1",
			},
			want: 1,
			want1: 7,

		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := convertToIntegers(nil, tt.args.to, tt.args.from)
			if got != tt.want {
				t.Errorf("convertToIntegers() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("convertToIntegers() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_computeForEachNumber(t *testing.T) {
	tests := []struct {
		name string
		number int
		want interface{}
	}{
		{
			name:           "when number is 15",
			number:         15,
			want:           "FizzBuzz",
		},
		{
			name:           "when number is 0",
			number:         0,
			want:           "Pink Flamingo",
		},
		{
			name:           "when number is 19",
			number:         19,
			want:           19,
		},
		{
			name:           "when number is 5",
			number:         5,
			want:           "Flamingo",
		},
		{
			name:           "when number is 3",
			number:         3,
			want:           "Flamingo",
		},
		{
			name:           "when number is 6",
			number:         6,
			want:           "Fizz",
		},
		{
			name:           "when number is 20",
			number:         20,
			want:           "Buzz",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := computeForEachNumber(tt.number); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("computeForEachNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}