package game

import (
	"reflect"
	"testing"
)

func TestNewSpellbook(t *testing.T) {
	tests := []struct {
		name string
		want Spellbook
	}{
		{
			name: "creating a new Spellbook",
			want: Spellbook{spells: make(map[string]uint64)},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSpellbook(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSpellbook() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSpellbook_IncreaseRandomSpell(t *testing.T) {

	testSpellBook2 := NewSpellbook()
	testSpellBook2.IncreaseRandomSpell(50, 50)

	type args struct {
		wisdom uint32
		level  uint64
	}
	tests := []struct {
		name string
		s    *Spellbook
		args args
	}{
		{
			name: "Increasing from a new spellbook",
			s:    createTestSpellBook(),
			args: args{
				wisdom: 15,
				level:  1,
			},
		},
		{
			name: "Using a spellbook with an existing spell",
			s:    &testSpellBook2,
			args: args{
				wisdom: 50,
				level:  50,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sb := copySpellBook(tt.s)
			tt.s.IncreaseRandomSpell(tt.args.wisdom, tt.args.level)
			if reflect.DeepEqual(*tt.s, sb) {
				t.Errorf("IncreaseRandomSpell() did not change any spells: %+v ", *tt.s)
			}
		})
	}
}

func createTestSpellBook() *Spellbook {
	sb := NewSpellbook()
	return &sb
}

func copySpellBook(sb *Spellbook) Spellbook {
	s := *sb

	newSpells := make(map[string]uint64)
	for k, v := range sb.spells {
		newSpells[k] = v
	}

	s.spells = newSpells

	return s
}
