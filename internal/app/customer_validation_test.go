package app

import (
	"testing"
)

func TestCustomerValidationExactMatch(t *testing.T) {
	service := CustomerValidationService{}
	match, _ := service.ValidateCustomer("John Doe", "John Doe", "john@example.com", "john@example.com", 0.8)

	if !match {
		t.Errorf("Expected exact match for name and email 'John Doe'")
	}
}

func TestCustomerValidationPhoneticMatch(t *testing.T) {
	service := CustomerValidationService{}
	match, _ := service.ValidateCustomer("Perez", "Peres", "perez@example.com", "peres@example.com", 0.8)

	if !match {
		t.Errorf("Expected phonetic match for 'Perez' and 'Peres'")
	}
}

func TestCustomerValidationPreventByronBrayan(t *testing.T) {
	service := CustomerValidationService{}
	match, _ := service.ValidateCustomer("Byron", "Brayan", "byron@example.com", "brayan@example.com", 0.8)

	if match {
		t.Errorf("Expected no match for 'Byron' and 'Brayan'")
	}
}

func TestCustomerValidationPreventPartialMatchBryanBrianne(t *testing.T) {
	service := CustomerValidationService{}
	match, _ := service.ValidateCustomer("Bryan", "Brianne", "bryan@example.com", "brianne@example.com", 0.8)

	if match {
		t.Errorf("Expected no match for 'Bryan' and 'Brianne'")
	}
}
