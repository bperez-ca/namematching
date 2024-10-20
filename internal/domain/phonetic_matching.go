package domain

import (
	"fmt"
	"github.com/dlclark/metaphone3"
)

// PhoneticMatch generates the Double Metaphone encoding for a name
func PhoneticMatch(name string) (string, string) {
	normalized := NormalizeName(name)
	mp := metaphone3.Encoder{}
	primaryKey, alternateKey := mp.Encode(normalized)
	fmt.Printf("Phonetic match for '%s' -> Primary: %s, Alternate: %s\n", name, primaryKey, alternateKey)
	return primaryKey, alternateKey
}
