package tdl

import (
	"math/rand"
	"testing"
	"time"
)

func TestSimpelWordGroups(t *testing.T) {
	for _, row := range tableSimpel {
		s := Distance(row.a, row.b)
		if s != row.score {
			t.Errorf("expected score == %d, got %d (a='%s', b='%s')", row.score, s, row.a, row.b)
		}
	}
}

var tableSimpel = []struct {
	score int
	a, b  string
}{
	{0, "", ""},             // empty string
	{0, "   ", ""},          // ignore white space in empty string
	{6, "azerty", ""},       // non empty against empty string
	{6, "", "qwerty"},       // empty against non empty string
	{2, "azer", "azerty"},   // adding to end
	{2, "erty", "azerty"},   // adding to start
	{2, "zert", "azerty"},   // adding to both ends
	{2, "azty", "azerty"},   // adding to middle
	{4, "zt", "azerty"},     // adding to middle and both ends
	{2, "azer", "azerty"},   // removing from end
	{2, "erty", "azerty"},   // removing from start
	{2, "zert", "azerty"},   // removing from both ends
	{2, "azty", "azerty"},   // removing from middle
	{4, "zt", "azerty"},     // removing from middle and both ends
	{2, "azErtY", "azerty"}, // substitution
	{1, "azrety", "azerty"}, // permutation

	// Couple of groups that's shown up in tests here and there,
	// doesn't hurt including them
	{4, "moral", "carry"},
	{5, "across", "is"},
	{4, "beak", "water"},
	{1, "teh", "the"},
	{1, "tets", "test"},
	{1, "fuor", "four"},
	{3, "kitten", "sitting"},
	{3, "Saturday", "Sunday"},
	{8, "rossettacode", "raisethysword"},
}

////////////////////////////////////////////////////////////////////////////////

func TestMetricSpaces(t *testing.T) {
	for _, f := range tableMetrics {
		if test := f(); test != true {
			t.Errorf("expected f() == true, got %t", test)
		}
	}
}

var tableMetrics = []func() bool{
	// metric space - identity of indiscernibles
	func() bool {
		a := "azerty"
		b := a
		return Distance(a, b) == 0
	},
	// metric space - symmetry
	func() bool {
		a := "azerty"
		b := "qwerty"
		return Distance(a, b) == Distance(b, a)
	},
	// metric space - triangle inequality edge 1
	func() bool {
		a := "rick"
		b := "rcik"
		c := "irkc"
		ab := Distance(a, b)
		ac := Distance(a, c)
		bc := Distance(b, c)
		return ab+ac >= bc
	},
	// metric space - triangle inequality edge 2
	func() bool {
		a := "rick"
		b := "rcik"
		c := "irkc"
		ab := Distance(a, b)
		ac := Distance(a, c)
		bc := Distance(b, c)
		return ab+bc >= ac
	},
	// metric space - triangle inequality edge 3
	func() bool {
		a := "rick"
		b := "rcik"
		c := "irkc"
		ab := Distance(a, b)
		ac := Distance(a, c)
		bc := Distance(b, c)
		return ac+bc >= ab
	},
	// metric space - triangle inequality quite / quiet / reject
	func() bool {
		a := "quite"
		b := "quiet"
		c := "reject"
		ab := Distance(a, b)
		ac := Distance(a, c)
		bc := Distance(b, c)
		eql1 := ab+ac >= bc
		eql2 := ab+bc >= ac
		eql3 := ac+bc >= ab
		return eql1 && eql2 && eql3
	},
	// metric space - triangle inequality leisure / Legislature / lies
	func() bool {
		a := "leisure"
		b := "Legislature"
		c := "lies"
		ab := Distance(a, b)
		ac := Distance(a, c)
		bc := Distance(b, c)
		eql1 := ab+ac >= bc
		eql2 := ab+bc >= ac
		eql3 := ac+bc >= ab
		return eql1 && eql2 && eql3
	},
}

////////////////////////////////////////////////////////////////////////////////

var seed = time.Now().Unix()

func init() {
	rand.Seed(seed)
}

func randWords(num int) []int {
	words := make([]int, num)
	for i := 0; i < num; i++ {
		words[i] = rand.Intn(len(wordsEnglish))
	}
	return words
}

// random words to check for non negativity
func TestRandomNonNegativity(t *testing.T) {
	ref := randWords(30)
	for i := 0; i < len(ref); i++ {
		for j := 0; j < len(wordsEnglish); j++ {
			a := Distance(wordsEnglish[ref[i]], wordsEnglish[j])
			positive := a >= 0
			if positive != true {
				t.Errorf("expected a >= 0, got %d (seed=%d, ref=#%d-%s, case=#%d-%s)",
					a, seed, ref[i], wordsEnglish[ref[i]], j, wordsEnglish[j],
				)
				break
			}
		}
	}
}

// random words to check for identity of indiscernibles
func TestRandomIndisceribles(t *testing.T) {
	ref := randWords(30)
	for i := 0; i < len(ref); i++ {
		for j := 0; j < len(wordsEnglish); j++ {
			a := Distance(wordsEnglish[ref[i]], wordsEnglish[j])
			var identity bool
			if a > 0 {
				identity = wordsEnglish[ref[i]] != wordsEnglish[j]
			} else {
				identity = wordsEnglish[ref[i]] == wordsEnglish[j]
			}
			if identity != true {
				t.Errorf("expected identity == true, got %t (seed=%d, ref=#%d-%s, case=#%d-%s)",
					identity, seed, ref[i], wordsEnglish[ref[i]], j, wordsEnglish[j],
				)
				break
			}
		}
	}
}

// random words to check for symmetry
func TestRandomSymmetry(t *testing.T) {
	ref := randWords(30)
	for i := 0; i < len(ref); i++ {
		for j := 0; j < len(wordsEnglish); j++ {
			a := Distance(wordsEnglish[ref[i]], wordsEnglish[j])
			b := Distance(wordsEnglish[j], wordsEnglish[ref[i]])
			symmetry := a == b
			if symmetry != true {
				t.Errorf("expected a == b, got a=%d, b=%d (seed=%d, ref=#%d-%s, case=#%d-%s)",
					a, b, seed, ref[i], wordsEnglish[ref[i]], j, wordsEnglish[j],
				)
				break
			}
		}
	}
}

// random words to check for triangle inequality
func TestRandomTriangleInequality(t *testing.T) {
	sets := 50
	triangles := 10000

	for i := 0; i < sets; i++ {
		for j := 0; j < triangles; j++ {
			// 3 randomly chosen words
			a := rand.Intn(len(wordsEnglish))
			b := rand.Intn(len(wordsEnglish))
			c := rand.Intn(len(wordsEnglish))
			ab := Distance(wordsEnglish[a], wordsEnglish[b])
			ac := Distance(wordsEnglish[a], wordsEnglish[c])
			bc := Distance(wordsEnglish[b], wordsEnglish[c])

			eql1 := ab+ac >= bc
			eql2 := ab+bc >= ac
			eql3 := ac+bc >= ab
			eql := eql1 && eql2 && eql3
			if eql != true {
				t.Errorf("expected eql == true, got %t (seed=%d, set=%d, a=%s, b=%s, c=%s)",
					eql, seed, i, wordsEnglish[a], wordsEnglish[b], wordsEnglish[c],
				)
				break
			}
		}
	}
}
