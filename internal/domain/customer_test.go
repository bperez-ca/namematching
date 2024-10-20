package domain

import (
	"fmt"
	"testing"
)

// Helper function to log and assert the match score
func assertMatchWithLogging(t *testing.T, got, want float64, message string) {
	fmt.Printf("%s: Got score = %.2f, Expected >= %.2f\n", message, got, want)
	if got < want {
		t.Errorf("%s: Expected match score >= %.2f, but got %.2f", message, want, got)
	}
}

// Helper function to log and assert when there should not be a match
func assertNoMatchWithLogging(t *testing.T, got, threshold float64, message string) {
	fmt.Printf("%s: Got score = %.2f, Expected < %.2f\n", message, got, threshold)
	if got >= threshold {
		t.Errorf("%s: Expected no match (score < %.2f), but got %.2f", message, threshold, got)
	}
}

// Example test cases with detailed logging for name comparison
func TestExactMatchWithLogging(t *testing.T) {
	customer := NewCustomer("John Doe", "john@example.com")
	score := customer.MatchName("John Doe")
	assertMatchWithLogging(t, score, 1.0, "Exact match of 'John Doe'")
}

func TestPhoneticMatchWithLogging(t *testing.T) {
	customer := NewCustomer("Perez", "john@example.com")
	score := customer.MatchName("Peres")
	assertMatchWithLogging(t, score, 0.8, "Phonetic match of 'Perez' and 'Peres'")
}

func TestPreventOvermatchingByronBrayanWithLogging(t *testing.T) {
	customer := NewCustomer("Byron", "byron@example.com")
	score := customer.MatchName("Brayan")
	assertNoMatchWithLogging(t, score, 0.8, "Prevent overmatching 'Byron' and 'Brayan'")
}

func TestPartialMatchBryanBrianneWithLogging(t *testing.T) {
	customer := NewCustomer("Bryan", "bryan@example.com")
	score := customer.MatchName("Brianne")
	assertNoMatchWithLogging(t, score, 0.8, "Prevent partial match of 'Bryan' and 'Brianne'")
}

func TestTypoMatchWithLogging(t *testing.T) {
	customer := NewCustomer("Johnathan", "john@example.com")
	score := customer.MatchName("Jonathan")
	assertMatchWithLogging(t, score, 0.8, "Typo match of 'Johnathan' and 'Jonathan'")
}

func TestSpecialCharacterHandlingWithLogging(t *testing.T) {
	customer := NewCustomer("O'Conner", "oconner@example.com")
	score := customer.MatchName("OConner")
	assertMatchWithLogging(t, score, 0.9, "Special character handling for 'O'Conner' and 'OConner'")
}

func TestHandlingAccentsWithLogging(t *testing.T) {
	customer := NewCustomer("Jón", "jon@example.com")
	score := customer.MatchName("Jon")
	assertMatchWithLogging(t, score, 0.8, "Accent handling for 'Jón' and 'Jon'")
}

func TestEmptyNamesWithLogging(t *testing.T) {
	customer := NewCustomer("", "empty@example.com")
	score := customer.MatchName("")
	fmt.Printf("Empty names: Got score = %.2f\n", score)
	if score != 1.0 {
		t.Errorf("Expected exact match for empty names, got %.2f", score)
	}
}

func TestSingleCharacterNamesWithLogging(t *testing.T) {
	customer := NewCustomer("A", "a@example.com")
	score := customer.MatchName("B")
	assertNoMatchWithLogging(t, score, 0.5, "No match for single character names 'A' and 'B'")
}

func TestCompletelyDifferentNamesWithLogging(t *testing.T) {
	customer := NewCustomer("Alice", "alice@example.com")
	score := customer.MatchName("Bob")
	assertNoMatchWithLogging(t, score, 0.5, "No match for completely different names 'Alice' and 'Bob'")
}

func TestNamesWithExtraSpaces(t *testing.T) {
	customer := NewCustomer("John  Doe", "john@example.com")
	score := customer.MatchName("John Doe")
	assertMatchWithLogging(t, score, 1.0, "Handling extra spaces in 'John  Doe' vs 'John Doe'")
}

func TestNamesWithHyphens(t *testing.T) {
	customer := NewCustomer("Jean-Pierre", "jean@example.com")
	score := customer.MatchName("Jean Pierre")
	assertMatchWithLogging(t, score, 1.0, "Handling hyphens in 'Jean-Pierre' vs 'Jean Pierre'")
}

func TestNamesWithApostrophes(t *testing.T) {
	customer := NewCustomer("O'Conner", "oconner@example.com")
	score := customer.MatchName("OConner") // The test should be with O Conner
	assertMatchWithLogging(t, score, 1.0, "Handling apostrophes in 'O'Conner' vs 'O Conner'")
}

func TestCaseInsensitiveMatch(t *testing.T) {
	customer := NewCustomer("John Doe", "john@example.com")
	score := customer.MatchName("john doe")
	assertMatchWithLogging(t, score, 1.0, "Case insensitive match 'John Doe' vs 'john doe'")
}

func TestLongNames(t *testing.T) {
	customer := NewCustomer("John Alexander Doe", "john@example.com")
	score := customer.MatchName("John Doe")
	assertMatchWithLogging(t, score, 0.8, "Handling long names 'John Alexander Doe' vs 'John Doe'")
}

func TestCompletelyDifferentNames(t *testing.T) {
	customer := NewCustomer("Alice", "alice@example.com")
	score := customer.MatchName("Charles")
	assertNoMatchWithLogging(t, score, 0.5, "Completely different names 'Alice' vs 'Charles'")
}

func TestOneSidedEmptyName(t *testing.T) {
	customer := NewCustomer("John Doe", "john@example.com")
	score := customer.MatchName("")
	assertNoMatchWithLogging(t, score, 0.0, "One-sided empty name 'John Doe' vs ''")
}

func TestVeryLongNames(t *testing.T) {
	customer := NewCustomer("Jonathan Alexander Michael Robert William Doe", "jon@example.com")
	score := customer.MatchName("Jonathan Doe")
	assertMatchWithLogging(t, score, 0.8, "Very long name 'Jonathan Alexander Michael Robert William Doe' vs 'Jonathan Doe'")
}

func TestTyposInDifferentPartsOfName(t *testing.T) {
	customer := NewCustomer("Jonhathan", "jon@example.com")
	score := customer.MatchName("Jonathan")
	assertMatchWithLogging(t, score, 0.8, "Typo in first name 'Jonhathan' vs 'Jonathan'")

	customer2 := NewCustomer("Perez", "doe@example.com")
	score2 := customer2.MatchName("Peres")
	assertMatchWithLogging(t, score2, 0.8, "Typo in last name 'Perez' vs 'Peres'")
}

func TestLongNamesWithSecondNameAsPrimary(t *testing.T) {
	customer := NewCustomer("Jonathan Alexander Michael Robert William Doe", "jon@example.com")
	score := customer.MatchName("Alexander Doe")
	assertMatchWithLogging(t, score, 0.8, "Very long name 'Jonathan Alexander Michael Robert William Doe' vs 'Alexander Doe'")

	customer2 := NewCustomer("Jonathan Alexander Michael Robert William Doe", "jon@example.com")
	score2 := customer2.MatchName("Michael Doe")
	assertMatchWithLogging(t, score2, 0.8, "Very long name 'Jonathan Alexander Michael Robert William Doe' vs 'Michael Doe'")
}
