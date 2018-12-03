package game

import "testing"

func Test_generateName(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "generate name",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := generateName()

			if len(got) == 0 {
				t.Errorf("generateName() = %v", got)
			}
		})
	}
}
