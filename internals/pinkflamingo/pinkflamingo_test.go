package pinkflamingo

import "testing"

func TestIsPinkFlamingo(t *testing.T) {
	tests := []struct {
		name string
		number int
		want bool
	}{
		{
			name:   "when number is 8",
			number: 8,
			want:   false,
		},
		{
			name:   "when number is 0",
			number: 0,
			want:   true,
		},
		{
			name:   "when number is 16",
			number: 16,
			want:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsPinkFlamingo(tt.number); got != tt.want {
				t.Errorf("IsPinkFlamingo() = %v, want %v", got, tt.want)
			}
		})
	}
}