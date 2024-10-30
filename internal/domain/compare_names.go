package domain

import (
	"fmt"
)

// CompareNames compares two names using tokenized comparison with a hybrid approach
func CompareNames(name1, name2 string) float64 {
	// Handle empty names explicitly
	if name1 == "" && name2 == "" {
		fmt.Printf("Both names are empty, returning perfect match score of 1.0\n")
		return 1.0
	}

	// Check for empty names and return a negative score for a no match
	if len(name1) == 0 || len(name2) == 0 {
		fmt.Printf("One of the names is empty, returning score -1.0\n")
		return -1.0
	}

	// Normalize both names
	normalized1 := NormalizeName(name1)
	normalized2 := NormalizeName(name2)

	// Step 1: Compare entire normalized names directly (to handle cases like "YukiMatsuda" vs "Yuki Matsuda")
	if normalized1 == normalized2 {
		fmt.Printf("Exact match for full names '%s' and '%s'\n", name1, name2)
		return 1.0
	}

	tokens1 := TokenizeName(name1)
	tokens2 := TokenizeName(name2)

	// Check if token slices are empty to prevent index out of range errors
	if len(tokens1) == 0 || len(tokens2) == 0 {
		fmt.Printf("One of the tokenized names is empty, returning score 0.0\n")
		return 0.0 // Handle empty token lists
	}

	totalScore := 0.0

	// Ensure we compare first and last names properly by treating middle names as optional
	firstName1 := tokens1[0]
	lastName1 := tokens1[len(tokens1)-1]

	firstName2 := tokens2[0]
	lastName2 := tokens2[len(tokens2)-1]

	firstNameScore, isFirstNameExactMatch := compareToken(firstName1, firstName2)

	lastNameScore, isLastNameExactMatch := compareToken(lastName1, lastName2)

	// If both first and last names are exact matches, treat it as a perfect match (score = 1.0)
	if isFirstNameExactMatch && isLastNameExactMatch {
		totalScore = 1.0
		fmt.Printf("Exact match for both first and last names, setting total score to 1.0\n")
	} else {
		// Names in between the first and the last name
		fullTokenScore := 0.0
		if len(tokens1) > 2 || len(tokens2) > 2 {
			for _, token1 := range tokens1 {
				bestScore := 0.0
				for _, token2 := range tokens2 {
					score, _ := compareToken(token1, token2)
					if bestScore < score {
						bestScore = score
					}
				}
				fullTokenScore += bestScore
			}
		}

		fmt.Printf("!!!---!!! Scores fist: '%.2f', last: '%.2f', fullToken: '%.2f'", firstNameScore, lastNameScore, fullTokenScore)
		// Total score is based on first name, last name, and middle name (if present)
		totalScore = firstNameScore + lastNameScore + fullTokenScore
	}

	fmt.Printf("Final total score: %.2f\n", totalScore)
	return totalScore
}

func compareToken(token1 string, token2 string) (float64, bool) {
	// Compare first names
	primary1, alternate1 := PhoneticMatch(token1)
	primary2, alternate2 := PhoneticMatch(token2)
	TokenScore := LevenshteinSimilarity(token1, token2)
	fmt.Printf("Comparing first names '%s' -> '%s', Phonetic: (%s, %s) vs (%s, %s)\n", token1, token2, primary1, alternate1, primary2, alternate2)

	isFirstNameExactMatch := false
	if primary1 == primary2 || alternate1 == alternate2 || primary1 == alternate2 || alternate1 == primary2 {
		if TokenScore >= 0.8 {
			TokenScore = 0.9
			isFirstNameExactMatch = true
		}
	}
	TokenScore *= 0.4 // Apply weight for first names
	fmt.Printf("Token score after weighting: %.2f\n", TokenScore)
	return TokenScore, isFirstNameExactMatch
}
