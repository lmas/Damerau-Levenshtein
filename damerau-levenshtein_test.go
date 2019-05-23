package tdl

import "testing"

var table = []struct {
	A     string
	B     string
	Score int
}{
	{"", "", 0},
	{" ", " ", 0},
	{"ABC", "ABC", 0},
	{"A", "B", 1},
	{"A", "", 1},
	{"", "B", 1},
	{"CA", "ABC", 2},
	{"aabc", "abc", 1},
	{"abcc", "abc", 1},
	{"abc", "bac", 1},
	{"abc", "abcc", 1},
	{"abc", "aabc", 1},
	{"abcdef", "poiu", 6},
	{"teh", "the", 1},
	{"tets", "test", 1},
	{"fuor", "four", 1},
	{"kitten", "sitting", 3},
	{"Saturday", "Sunday", 3},
	{"rosettacode", "raisethysword", 8},
}

func TestTable(t *testing.T) {
	for _, row := range table {
		score := Distance(row.A, row.B)
		if score != row.Score {
			t.Errorf("Expected score=%d, got score=%d (A='%s', B='%s')", row.Score, score, row.A, row.B)
		}
	}
}
