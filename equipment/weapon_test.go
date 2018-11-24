package equipment

import (
	"testing"
)

func Test_generateWeapon(t *testing.T) {
	type args struct {
		level int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test level 0",
			args: args{level: 0},
			want: 0,
		},
		{
			name: "test level -1",
			args: args{level: -1},
			want: 0,
		},
		{
			name: "test level 10",
			args: args{level: 10},
			want: 10,
		},
		{
			name: "Level above the number of base weapons",
			args: args{level: len(weaponBase) + 1},
			want: len(weaponBase) + 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := generateWeapon(tt.args.level)
			if got.Name == "" {
				t.Errorf("generateWeapon() Name field is empty and shouldn't be")
			}
			if (got.Bonus + got.Level) > tt.want {
				t.Errorf("generateWeapon() total weapon level = %d, but should not be over %d", got.Bonus+got.Level, tt.want)
			}
		})
	}
}
