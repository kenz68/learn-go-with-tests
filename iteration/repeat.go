package iteration

import "strings"

// Repeat returns a string with the character repeated n times
func Repeat(character string, n int) string {
	var repeated strings.Builder // its perfromance is better than using string concatenation from 136 ns/op to 23.91 ns/op
	for i := 0; i < n; i++ {
		repeated.WriteString(character)
	}
	return repeated.String()
}
