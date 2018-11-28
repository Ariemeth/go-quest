package comparison

import "testing"

func TestMinInt2(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "two positive values",
			args: args{a: 1, b: 10},
			want: 1,
		},
		{
			name: "two positive values with values swapped",
			args: args{a: 10, b: 1},
			want: 1,
		},
		{
			name: "two negative values",
			args: args{a: -1, b: -10},
			want: -10,
		},
		{
			name: "two negative values with values swapped",
			args: args{a: -10, b: -1},
			want: -10,
		},
		{
			name: "one negative and 1 positive value",
			args: args{a: 1, b: -10},
			want: -10,
		},
		{
			name: "one negative and 1 positive value values swapped",
			args: args{a: -10, b: 1},
			want: -10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MinInt2(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("MinInt2() = %v, want %v", got, tt.want)
			}
		})
	}
}
