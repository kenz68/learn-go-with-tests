package iteration

import "strings"

const repeatCount = 5

// Repeat returns a string with the character repeated 5 times
func Repeat(character string) string {
	var repeated strings.Builder // its perfromance is better than using string concatenation from 136 ns/op to 23.91 ns/op
	for i := 0; i < repeatCount; i++ {
		repeated.WriteString(character)
	}
	return repeated.String()
}
