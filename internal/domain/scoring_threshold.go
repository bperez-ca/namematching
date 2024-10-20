package domain

// IsMatch checks if the similarity score exceeds the threshold
func IsMatch(score float64, threshold float64) bool {
	return score >= threshold
}
