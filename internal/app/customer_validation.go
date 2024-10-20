package app

import "NameMatching/internal/domain"

// CustomerValidationService orchestrates customer validation (use case)
type CustomerValidationService struct{}

// ValidateCustomer orchestrates the validation of two customers' names and emails
func (s *CustomerValidationService) ValidateCustomer(name1, name2, email1, email2 string, threshold float64) (bool, float64) {
	// Create customer domain objects
	customer1 := domain.NewCustomer(name1, email1)
	customer2 := domain.NewCustomer(name2, email2)

	// Perform name and email matching using domain logic
	nameScore := customer1.MatchName(customer2.Name)
	emailScore := customer1.MatchEmail(customer2.Email)

	// Combine the scores (you can refine this logic as per your business need)
	finalScore := (nameScore + emailScore) / 2

	// Apply the threshold check
	return domain.IsMatch(finalScore, threshold), finalScore
}
