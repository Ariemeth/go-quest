package equipment

import (
	"testing"
)

func Test_attributes_contains(t *testing.T) {
	type args struct {
		att attribute
	}
	tests := []struct {
		name string
		a    attributes
		args args
		want bool
	}{
		{
			name: "item in slice",
			a:    offenseAttrib,
			args: args{att: offenseAttrib[0]},
			want: true,
		},
		{
			name: "item not in slice",
			a:    offenseAttrib,
			args: args{att: attribute{name: "should not be in slice"}},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.contains(tt.args.att); got != tt.want {
				t.Errorf("attributes.contains() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_absInt(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "positive value",
			args: args{10},
			want: 10,
		},
		{
			name: "negative value",
			args: args{-10},
			want: 10,
		},
		{
			name: "zero",
			args: args{0},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := absInt(tt.args.x); got != tt.want {
				t.Errorf("absInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
