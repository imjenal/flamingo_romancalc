package romancalculator

import (
	"reflect"
	"testing"
)

func TestEvaluateinBodmas(t *testing.T) {
	tests := []struct {
		expr    string
		want    interface{}
		wantErr error
	}{
		{
			expr:    "6 + 7 * 2",
			want:    float64(20),
			wantErr: nil,
		},
		{
			expr:    "6 + 7 - 2",
			want:    float64(11),
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		got, err := EvaluateinBodmas(tt.expr)
		if err != tt.wantErr {
			t.Errorf("EvaluateinBodmas() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
		if !reflect.DeepEqual(got, tt.want) {

			t.Errorf("EvaluateinBodmas() got = %v, want %v", got, tt.want)
		}
	}
}

func TestIsOperator(t *testing.T) {
	tests := []struct {
		name string
		op   string
		want bool
	}{
		{
			name: "when op is (",
			op:   "(",
			want: true,
		},
		{
			name: "when op is tsh",
			op:   "tsh",
			want: false,
		},
		{
			name: "when op is *",
			op:   "*",
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsOperator(tt.op); got != tt.want {
				t.Errorf("IsOperator() = %v, want %v", got, tt.want)
			}
		})
	}
}
