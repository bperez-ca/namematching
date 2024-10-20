package domain

import (
	"github.com/agnivade/levenshtein"
)

// LevenshteinSimilarity calculates the Levenshtein similarity score between two names
func LevenshteinSimilarity(name1, name2 string) float64 {
	name1 = NormalizeName(name1)
	name2 = NormalizeName(name2)

	dist := levenshtein.ComputeDistance(name1, name2)
	maxLen := float64(maxIntegers(len(name1), len(name2)))
	if maxLen == 0 {
		return 1.0
	}
	return 1.0 - float64(dist)/maxLen
}

// max returns the maximum of two integers
func maxIntegers(a, b int) int {
	if a > b {
		return a
	}
	return b
}
