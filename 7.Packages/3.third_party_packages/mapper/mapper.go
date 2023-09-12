package mapper

import (
	"github.com/rylans/getlang"
)

// Greet takes in a string as an argument
// It detects your language from the given string
// And greets back in that language
func Greet(s string) string {
	info := getlang.FromString(s)
	switch info.LanguageCode() {
	case "en":
		return "Hello!"
	case "fr":
		return "Bonjour!"
	case "de":
		return "Guten Tag!"
	default:
		return "I don't know your language yet!"
	}
}
