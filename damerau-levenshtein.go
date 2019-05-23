package tdl

import "strings"

// Reference:
// https://en.wikipedia.org/wiki/Damerau%E2%80%93Levenshtein_distance#Distance_with_adjacent_transpositions

// Return the smalles int from a list
func minimum(is ...int) int {
	min := is[0]
	for _, i := range is {
		if min > i {
			min = i
		}
	}
	return min
}

// Calculate and return the true Damerau–Levenshtein distance of string A and B.
func Distance(a, b string) int {
	a, b = strings.TrimSpace(a), strings.TrimSpace(b)
	lenA, lenB := len(a), len(b)
	if lenA == 0 {
		if lenB == 0 {
			return 0
		}
		return lenB
	} else if lenB == 0 {
		return lenA
	}

	da := make(map[rune]int)
	for _, rn := range a + b {
		da[rn] = 0
	}

	d := make([][]int, lenA+2)
	for i := range d {
		d[i] = make([]int, lenB+2)
	}

	maxDist := lenA + lenB
	d[0][0] = maxDist
	for i := 0; i <= lenA; i++ {
		d[i+1][0] = maxDist
		d[i+1][1] = i
	}
	for j := 0; j <= lenB; j++ {
		d[0][j+1] = maxDist
		d[1][j+1] = j
	}

	for i := 1; i <= lenA; i++ {
		db := 0
		for j := 1; j <= lenB; j++ {
			k := da[rune(b[j-1])]
			l := db
			cost := 1
			if a[i-1] == b[j-1] {
				cost = 0
				db = j
			}

			d[i+1][j+1] = minimum(
				d[i][j]+cost,              // substitution
				d[i+1][j]+1,               // insertion
				d[i][j+1]+1,               // deletion
				d[k][l]+(i-k-1)+1+(j-l-1), // transposition
			)
		}
		da[rune(a[i-1])] = i
	}
	return d[lenA+1][lenB+1]
}
