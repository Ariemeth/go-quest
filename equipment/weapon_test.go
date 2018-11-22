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
	}{
		{
			name: "test level 0",
			args: args{level: 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := generateWeapon(tt.args.level)
			if got.Name == "" {
				t.Errorf("generateWeapon() Name field is empty and shouldn't be")
			}
		})
	}
}
