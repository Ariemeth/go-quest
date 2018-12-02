package main

import (
	"testing"
)

func Test_stats_Total(t *testing.T) {
	tests := []struct {
		name string
		s    stats
		want uint
	}{
		{
			name: "total value",
			s:    stats{10, 10, 10, 10, 10, 10},
			want: 60,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Total(); got != tt.want {
				t.Errorf("stats.Total() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_generateInitialStats(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "generate initial stats",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := generateInitialStats()

			if got.Charisma <= 0 {
				t.Errorf("Charisma = %v, should be positive", got.Charisma)
			}
			if got.Constitution <= 0 {
				t.Errorf("Constitution = %v, should be positive", got.Constitution)
			}
			if got.Dexterity <= 0 {
				t.Errorf("Intelligence = %v, should be positive", got.Intelligence)
			}
			if got.Intelligence <= 0 {
				t.Errorf("Charisma = %v, should be positive", got.Charisma)
			}
			if got.Strength <= 0 {
				t.Errorf("Strength = %v, should be positive", got.Strength)
			}
			if got.Wisdom <= 0 {
				t.Errorf("Wisdom = %v, should be positive", got.Wisdom)
			}
		})
	}
}
