// Package tdl implements the true Damerau–Levenshtein distance.
//
// Reference:
// https://en.wikipedia.org/wiki/Damerau%E2%80%93Levenshtein_distance#Distance_with_adjacent_transpositions
package tdl

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

// Distance calculates and returns the true Damerau–Levenshtein distance of string A and B.
// It's the caller's responsibility if he wants to trim whitespace or fix lower/upper cases.
func Distance(a, b string) int {
	lenA, lenB := len(a), len(b)
	if lenA < 1 {
		return lenB
	}
	if lenB < 1 {
		return lenA
	}

	matrix := make([][]int, lenA+2)
	for i := range matrix {
		matrix[i] = make([]int, lenB+2)
	}
	matrix[0][0] = lenA + lenB + 1
	for i := 0; i <= lenA; i++ {
		matrix[i+1][1] = i
		matrix[i+1][0] = matrix[0][0]
	}
	for j := 0; j <= lenB; j++ {
		matrix[1][j+1] = j
		matrix[0][j+1] = matrix[0][0]
	}

	da := make(map[rune]int)
	for _, r := range a + b {
		da[r] = 0
	}

	for i := 1; i <= lenA; i++ {
		db := 0
		for j := 1; j <= lenB; j++ {
			i1 := da[rune(b[j-1])]
			j1 := db
			cost := 1
			if a[i-1] == b[j-1] {
				cost = 0
				db = j
			}

			// By "conventional wisdom", the costs for the ins/del/trans operations are always +1
			matrix[i+1][j+1] = minimum(
				matrix[i][j]+cost, // substitution
				matrix[i+1][j]+1,  // insertion
				matrix[i][j+1]+1,  // deletion
				matrix[i1][j1]+(i-i1-1)+1+(j-j1-1), // transposition
			)
		}
		da[rune(a[i-1])] = i
	}
	return matrix[lenA+1][lenB+1]
}
