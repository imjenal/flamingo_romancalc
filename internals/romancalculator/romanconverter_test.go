package romancalculator

import "testing"

func TestToRoman(t *testing.T) {
	tests := []struct {
		name string
		number int
		roman string
	}{
		{
			name:"when number is 0",
			number:0,
			roman: "",
		},
		{
			name:"when number is 1",
			number:1,
			roman: "I",
		},
		{
			name:"when number is 1993",
			number:1993,
			roman: "MCMXCIII",
		},
		{
			name:"when number is 666",
			number:666,
			roman: "DCLXVI",
		},
		{
			name:"when number is 444",
			number:444,
			roman: "CDXLIV",
		},
		{
			name:"when number is 555",
			number:555,
			roman: "DLV",
		},
		{
			name:"when number is 999",
			number:999,
			roman: "CMXCIX",
		},
		{
			name:"when number is 555",
			number:555,
			roman: "DLV",
		},
		{
			name:"when number is 1111",
			number:1111,
			roman: "MCXI",
		},
		{
			name:"when number is 2222",
			number:2222,
			roman: "MMCCXXII",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			roman := NewRoman()
			if got := roman.ToRoman(tt.number); got != tt.roman {
				t.Errorf("FromInt() = %v, want %v", got, tt.roman)
			}
		})
	}
}

func TestRoman_ToNumber(t *testing.T) {
	tests := []struct {
		name string
		roman string
		want int
	}{
		{
			name:"when number is 0",
			roman:"",
			want: 0,
		},
		{
			name:"when number is 1",
			roman:"I",
			want: 1,
		},
		{
			name:"when number is 1993",
			roman:"MCMXCIII",
			want: 1993,
		},
		{
			name:"when number is 666",
			roman:"DCLXVI",
			want: 666,
		},
		{
			name:"when number is 444",
			roman:"CDXLIV",
			want: 444,
		},
		{
			name:"when number is 555",
			roman:"DLV",
			want: 555,
		},
		{
			name:"when number is 999",
			roman:"CMXCIX",
			want: 999,
		},
		{
			name:"when number is 1111",
			roman:"MCXI",
			want: 1111,
		},
		{
			name:"when number is 2222",
			roman:"MMCCXXII",
			want: 2222,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Roman{}
			if got := r.ToNumber(tt.roman); got != tt.want {
				t.Errorf("ToNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}