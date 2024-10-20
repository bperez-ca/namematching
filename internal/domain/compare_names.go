package domain

import "fmt"

// CompareNames compares two names using tokenized comparison with a hybrid approach
func CompareNames(name1, name2 string) float64 {
	// Handle empty names explicitly
	if name1 == "" && name2 == "" {
		fmt.Printf("Both names are empty, returning perfect match score of 1.0\n")
		return 1.0
	}

	tokens1 := TokenizeName(name1)
	tokens2 := TokenizeName(name2)

	totalScore := 0.0
	maxTokens := maxIntegers(len(tokens1), len(tokens2)) // maxTokens is an integer now

	for i := 0; i < maxTokens; i++ { // Loop works with integers
		var token1, token2 string

		if i < len(tokens1) {
			token1 = tokens1[i]
		}
		if i < len(tokens2) {
			token2 = tokens2[i]
		}

		// Skip comparison if one of the tokens is empty (mismatched number of tokens)
		if token1 == "" || token2 == "" {
			continue
		}

		primary1, alternate1 := PhoneticMatch(token1)
		primary2, alternate2 := PhoneticMatch(token2)

		fmt.Printf("Comparing '%s' -> '%s', Phonetic: (%s, %s) vs (%s, %s)\n", token1, token2, primary1, alternate1, primary2, alternate2)

		tokenScore := LevenshteinSimilarity(token1, token2)
		fmt.Printf("Levenshtein score for '%s' vs '%s' = %.2f\n", token1, token2, tokenScore)

		if primary1 == primary2 || alternate1 == alternate2 || primary1 == alternate2 || alternate1 == primary2 {
			if tokenScore >= 0.8 {
				tokenScore = 1.0
			}
		}

		totalScore += tokenScore
	}

	finalScore := totalScore / float64(maxTokens)
	fmt.Printf("Final name match score for '%s' vs '%s' = %.2f\n", name1, name2, finalScore)
	return finalScore
}
