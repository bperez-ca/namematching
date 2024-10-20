package domain

import "fmt"

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

	tokens1 := TokenizeName(name1)
	tokens2 := TokenizeName(name2)

	totalScore := 0.0

	// Ensure we compare first and last names properly by treating middle names as optional
	firstName1 := tokens1[0]
	lastName1 := tokens1[len(tokens1)-1]

	firstName2 := tokens2[0]
	lastName2 := tokens2[len(tokens2)-1]

	// Compare first names
	primary1, alternate1 := PhoneticMatch(firstName1)
	primary2, alternate2 := PhoneticMatch(firstName2)
	firstNameScore := LevenshteinSimilarity(firstName1, firstName2)
	fmt.Printf("Comparing first names '%s' -> '%s', Phonetic: (%s, %s) vs (%s, %s)\n", firstName1, firstName2, primary1, alternate1, primary2, alternate2)

	isFirstNameExactMatch := false
	if primary1 == primary2 || alternate1 == alternate2 || primary1 == alternate2 || alternate1 == primary2 {
		if firstNameScore >= 0.8 {
			firstNameScore = 0.9
			isFirstNameExactMatch = true
		}
	}
	firstNameScore *= 0.4 // Apply weight for first names
	fmt.Printf("First name score after weighting: %.2f\n", firstNameScore)

	// Compare last names
	primary1, alternate1 = PhoneticMatch(lastName1)
	primary2, alternate2 = PhoneticMatch(lastName2)
	lastNameScore := LevenshteinSimilarity(lastName1, lastName2)
	fmt.Printf("Comparing last names '%s' -> '%s', Phonetic: (%s, %s) vs (%s, %s)\n", lastName1, lastName2, primary1, alternate1, primary2, alternate2)

	isLastNameExactMatch := false
	if primary1 == primary2 || alternate1 == alternate2 || primary1 == alternate2 || alternate1 == primary2 {
		if lastNameScore >= 0.8 {
			lastNameScore = 0.9
			isLastNameExactMatch = true
		}
	}
	lastNameScore *= 0.4 // Apply weight for last names
	fmt.Printf("Last name score after weighting: %.2f\n", lastNameScore)

	// If both first and last names are exact matches, treat it as a perfect match (score = 1.0)
	if isFirstNameExactMatch && isLastNameExactMatch {
		totalScore = 1.0
		fmt.Printf("Exact match for both first and last names, setting total score to 1.0\n")
	} else {
		// Middle name handling (only if both names have middle names)
		middleNameScore := 0.0
		if len(tokens1) > 2 && len(tokens2) > 2 { // Both names have a middle name
			middleName1 := tokens1[1]
			middleName2 := tokens2[1]

			primary1, alternate1 = PhoneticMatch(middleName1)
			primary2, alternate2 = PhoneticMatch(middleName2)
			middleNameScore = LevenshteinSimilarity(middleName1, middleName2)
			fmt.Printf("Comparing middle names '%s' -> '%s', Phonetic: (%s, %s) vs (%s, %s)\n", middleName1, middleName2, primary1, alternate1, primary2, alternate2)

			if primary1 == primary2 || alternate1 == alternate2 || primary1 == alternate2 || alternate1 == primary2 {
				if middleNameScore >= 0.8 { // Allow slight differences
					middleNameScore = 0.9 // Apply smaller penalty
				}
			}
			middleNameScore *= 0.2 // Apply lower weight for middle names
			fmt.Printf("Middle name score after weighting: %.2f\n", middleNameScore)
		}

		// Total score is based on first name, last name, and middle name (if present)
		totalScore = firstNameScore + lastNameScore + middleNameScore
	}

	fmt.Printf("Final total score: %.2f\n", totalScore)
	return totalScore
}
