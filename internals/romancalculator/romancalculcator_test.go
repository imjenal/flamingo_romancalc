package romancalculator

import (
	"testing"
)

func Test_numberEquivalentExpression(t *testing.T) {
	tests := []struct {
		name string
		expr string
		want string
	}{
		{
			name: "when expression is X * VII",
			expr: "X * VII",
			want: "10*7",
		},
		{
			name: "when expression is (X * III / IV)",
			expr: "(X * III / IV)",
			want: "10*3/4",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numberEquivalentExpression(tt.expr, NewRoman()); got != tt.want {
				t.Errorf("numberEquivalentExpression() = %v, want %v", got, tt.want)
			}
		})
	}
}
