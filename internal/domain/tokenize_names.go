package domain

import (
	"regexp"
	"strings"
)

func TokenizeName(name string) []string {
	// Normalize the name first (e.g., lowercase and remove diacritics)
	normalized := NormalizeName(name)

	// Replace hyphens and apostrophes with spaces
	normalized = strings.ReplaceAll(normalized, "-", " ")
	normalized = strings.ReplaceAll(normalized, "'", " ")
	//fmt.Printf("Normalized Name '%s' -> '%s',\n", name, normalized)
	// Use a regular expression to split based on spaces (including multiple spaces)
	re := regexp.MustCompile(`\s+`)
	tokens := re.Split(normalized, -1)
	//fmt.Printf("Tokens Name '%s' -> '%s',\n", name, tokens)
	return tokens
}
