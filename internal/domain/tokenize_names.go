package domain

import (
	"regexp"
	"strings"
)

// TokenizeName splits a name into tokens, replaces special characters with spaces, removes accents,
// and ensures case-insensitive comparison.
func TokenizeName(name string) []string {
	// Replace hyphens and apostrophes with spaces
	name = strings.ReplaceAll(name, "-", " ")
	name = strings.ReplaceAll(name, "'", " ")

	// Remove accents and convert to lowercase for case-insensitive matching
	name = NormalizeName(name)

	// Replace all non-alphanumeric characters with spaces
	re := regexp.MustCompile(`[^\w]`)
	cleanedName := re.ReplaceAllString(name, " ")

	// Split the cleaned name into tokens by spaces
	tokens := strings.Fields(cleanedName)

	return tokens
}
