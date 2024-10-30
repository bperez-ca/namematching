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

// NormalizeName standardizes the name by converting it to lowercase, removing diacritics,
// and replacing special characters (except hyphens) with spaces. Apostrophes are removed and joined.
func NormalizeName(name string) string {
	// Convert to lowercase and remove diacritics
	normalized := strings.ToLower(RemoveDiacritics(name))

	// Replace all special characters (except hyphens) with spaces
	var sb strings.Builder
	for _, r := range normalized {
		if unicode.IsLetter(r) || unicode.IsSpace(r) || r == '-' {
			sb.WriteRune(r) // Keep letters, spaces, and hyphens
		} else {
			sb.WriteRune(' ') // Replace other special characters with spaces
		}
	}

	return strings.TrimSpace(sb.String())
}
