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

func TestSpecialCharacterHandlingWithAllCases(t *testing.T) {
	// Test case for "O'Conner" and "OConner"
	customer := NewCustomer("O'Conner", "oconner@example.com")
	score := customer.MatchName("OConner")
	assertMatchWithLogging(t, score, 0.9, "Special character handling for 'O'Conner' and 'OConner'")

	// Test case for "O'Conner" and "O Conner"
	customer2 := NewCustomer("O'Conner", "oconner@example.com")
	score2 := customer2.MatchName("O Conner")
	assertMatchWithLogging(t, score2, 0.9, "Special character handling for 'O'Conner' and 'O Conner'")

	// Test case for "OConner" and "O Conner"
	customer3 := NewCustomer("OConner", "oconner@example.com")
	score3 := customer3.MatchName("O Conner")
	assertMatchWithLogging(t, score3, 0.9, "Special character handling for 'OConner' and 'O Conner'")
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

func TestLongNamesWithPreferredNames(t *testing.T) {
	customer := NewCustomer("Brayan Ferney Perez Moreno", "jon@example.com")
	score := customer.MatchName("Brayan Perez")
	assertMatchWithLogging(t, score, 0.8, "Very long name 'Brayan Ferney Perez Moreno' vs 'Brayan Perez'")

	customer2 := NewCustomer("Brayan Ferney Perez Moreno", "jon@example.com")
	score2 := customer2.MatchName("Ferney Perez")
	assertMatchWithLogging(t, score2, 0.8, "Very long name 'Brayan Ferney Perez Moreno' vs 'Ferney Perez'")

	customer3 := NewCustomer("Brayan Ferney Perez Moreno", "jon@example.com")
	score3 := customer3.MatchName("Ferney Perez Moreno")
	assertMatchWithLogging(t, score3, 0.8, "Very long name 'Brayan Ferney Perez Moreno' vs 'Ferney Perez Moreno'")

	customer4 := NewCustomer("Brayan Ferney Perez Moreno", "jon@example.com")
	score4 := customer4.MatchName("Brayan F. Perez")
	assertMatchWithLogging(t, score4, 0.8, "Very long name 'Brayan Ferney Perez Moreno' vs 'Brayan F. Perez'")

	customer5 := NewCustomer("Brayan Ferney Perez Moreno", "jon@example.com")
	score5 := customer5.MatchName("Brayan Perez Moreno")
	assertMatchWithLogging(t, score5, 0.8, "Very long name 'Brayan Ferney Perez Moreno' vs 'Brayan Perez Moreno'")

	customer6 := NewCustomer("Brayan Ferney Perez-Moreno", "jon@example.com")
	score6 := customer6.MatchName("Brayan Perez Moreno")
	assertMatchWithLogging(t, score6, 0.8, "Very long name 'Brayan Ferney Perez-Moreno' vs 'Brayan Perez Moreno'")
}

func TestEmptyNameOrSpecialCharacters(t *testing.T) {
	customer := NewCustomer("!!!", "brayan@example.com")
	score := customer.MatchName("Brayan Perez")
	assertMatchWithLogging(t, score, 0.0, "Special character-only name '!!!' vs 'Brayan Perez'")
}

func TestLongNameWithSpecialCharacters(t *testing.T) {
	customer7 := NewCustomer("Brayan+Ferney Perez+Moreno", "jon@example.com")
	score7 := customer7.MatchName("Brayan Perez")
	assertMatchWithLogging(t, score7, 0.8, "Very long name 'Brayan+Ferney Perez+Moreno' vs 'Brayan Perez'")
}

func TestLongNameNotMatchingPhoneticallySimilar(t *testing.T) {
	customer := NewCustomer("Byron Fernando Piedrahita Moreno", "jon@example.com")
	score := customer.MatchName("Brayan Perez")
	assertNoMatchWithLogging(t, score, 0.8, "Very long name 'Byron Fernando Piedrahita Moreno' vs 'Brayan Perez'")

	customer2 := NewCustomer("Brian Piers", "jon@example.com")
	score2 := customer2.MatchName("Brayan Perez")
	assertNoMatchWithLogging(t, score2, 0.8, "Very long name 'Brian Piers' vs 'Brayan Perez'")

	customer3 := NewCustomer("Bryan Paris", "jon@example.com")
	score3 := customer3.MatchName("Brayan Perez")
	assertNoMatchWithLogging(t, score3, 0.8, "Very long name 'Bryan Paris' vs 'Brayan Perez'")
}

func TestNameWithPercentageSign(t *testing.T) {
	customer := NewCustomer("Brayan%Perez", "brayan@example.com")
	score := customer.MatchName("Brayan Perez")
	assertMatchWithLogging(t, score, 0.8, "Name with percentage sign 'Brayan%Perez' vs 'Brayan Perez'")
}

func TestNameWithDollarSign(t *testing.T) {
	customer := NewCustomer("Brayan$Perez", "brayan@example.com")
	score := customer.MatchName("Brayan Perez")
	assertMatchWithLogging(t, score, 0.8, "Name with dollar sign 'Brayan$Perez' vs 'Brayan Perez'")
}

func TestLongNameWithVariousSpecialCharacters(t *testing.T) {
	customer := NewCustomer("Brayan@#*!|~Ferney%$Perez@#Moreno", "brayan@example.com")
	score := customer.MatchName("Brayan Perez")
	assertMatchWithLogging(t, score, 0.8, "Very long name 'Brayan@#*!|~Ferney%$Perez@#Moreno' vs 'Brayan Perez'")
}

func TestEthnicNamesWithSpecialCharacters(t *testing.T) {
	// Test case for Indian name with spaces and special characters
	customer := NewCustomer("Rajesh Kumar", "rajesh.kumar@example.com")
	score := customer.MatchName("Rajesh Kumar")
	assertMatchWithLogging(t, score, 1.0, "Exact match for Indian name 'Rajesh Kumar'")

	// Test case for Indian name with missing last name
	customer2 := NewCustomer("Rajesh Kumar", "rajesh.kumar@example.com")
	score2 := customer2.MatchName("Rajesh")
	assertMatchWithLogging(t, score2, 0.8, "Indian name 'Rajesh Kumar' vs 'Rajesh'")

	// Test case for Japanese name with space between first and last name
	customer3 := NewCustomer("Yuki Matsuda", "yuki.matsuda@example.jp")
	score3 := customer3.MatchName("Yuki Matsuda")
	assertMatchWithLogging(t, score3, 1.0, "Exact match for Japanese name 'Yuki Matsuda'")

	// Test case for Japanese name without space
	customer4 := NewCustomer("YukiMatsuda", "yuki.matsuda@example.jp")
	score4 := customer4.MatchName("Yuki Matsuda")
	assertMatchWithLogging(t, score4, 0.9, "Japanese name 'YukiMatsuda' vs 'Yuki Matsuda'")

	// Test case for Brazilian name with accent
	customer5 := NewCustomer("José da Silva", "jose.silva@example.com")
	score5 := customer5.MatchName("Jose da Silva")
	assertMatchWithLogging(t, score5, 1.0, "Brazilian name 'José da Silva' vs 'Jose da Silva'")

	// Test case for Brazilian name with missing first name
	customer6 := NewCustomer("José da Silva", "jose.silva@example.com")
	score6 := customer6.MatchName("da Silva")
	assertMatchWithLogging(t, score6, 0.8, "Brazilian name 'José da Silva' vs 'da Silva'")

	// Test case for Brazilian name with special characters
	customer7 := NewCustomer("Ana-Maria", "ana.maria@example.com")
	score7 := customer7.MatchName("Ana Maria")
	assertMatchWithLogging(t, score7, 0.9, "Brazilian name with hyphen 'Ana-Maria' vs 'Ana Maria'")

	// Test case for Japanese name with special characters
	customer8 := NewCustomer("Takashi O’Nishi", "takashi.nishi@example.jp")
	score8 := customer8.MatchName("Takashi ONishi")
	assertMatchWithLogging(t, score8, 0.9, "Japanese name with apostrophe 'Takashi O’Nishi' vs 'Takashi ONishi'")
}
