package domain

// Customer represents a customer entity in the system
type Customer struct {
	Name  string
	Email string
}

// NewCustomer creates a new Customer instance
func NewCustomer(name, email string) *Customer {
	return &Customer{Name: name, Email: email}
}

// MatchName compares two names using tokenized comparison
func (c *Customer) MatchName(otherName string) float64 {
	return CompareNames(c.Name, otherName)
}

// MatchEmail compares two emails using Levenshtein similarity
func (c *Customer) MatchEmail(otherEmail string) float64 {
	return LevenshteinSimilarity(c.Email, otherEmail)
}
