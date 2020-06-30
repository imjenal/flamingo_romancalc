package romancalculator

import "testing"

func TestFromInt(t *testing.T) {
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
			if got := FromInt(tt.number); got != tt.roman {
				t.Errorf("FromInt() = %v, want %v", got, tt.roman)
			}
		})
	}
}