package anagram

import (
	"testing"
)

func Test_Verify(t *testing.T) {
	type args struct {
		x string
		y string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Word: Correct", args{"trings", "string"}, true},
		{"Word: Incorrect", args{"trings", "hello"}, false},
		{"Phrase: Correct", args{"Dormitory", "dirty room"}, true},
		{"Phrase: Incorrect", args{"This is a Doctor", "dirty room"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Verify(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("IsAnagram() = %v, want %v", got, tt.want)
			}
		})
	}
}
