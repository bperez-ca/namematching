package domain

import (
	"golang.org/x/text/runes"
	"golang.org/x/text/unicode/norm"
	"strings"
	"unicode"
)

// RemoveDiacritics removes accents (diacritical marks) from a string
func RemoveDiacritics(input string) string {
	// Normalize to decomposed form
	t := norm.NFD.String(input)

	// Remove all non-spacing marks (diacritics)
	t = runes.Remove(runes.In(unicode.Mn)).String(t)

	return t
}

// NormalizeName standardizes the name by converting it to lowercase, removing diacritics, and preserving hyphens
func NormalizeName(name string) string {
	// Convert to lowercase and remove diacritics
	normalized := strings.ToLower(RemoveDiacritics(name))

	// Replace apostrophes with spaces (to handle cases like "O'Conner")
	normalized = strings.ReplaceAll(normalized, "'", "")

	// Keep hyphens, but remove other punctuation like apostrophes
	var sb strings.Builder
	for _, r := range normalized {
		if unicode.IsLetter(r) || r == '-' || unicode.IsSpace(r) {
			sb.WriteRune(r)
		}
	}

	return strings.TrimSpace(sb.String())
}
