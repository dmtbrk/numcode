package numcode

import "strings"

var encodeTable = map[rune]rune{'a': '1', 'e': '2', 'i': '3', 'o': '4', 'u': '5', 'A': '1', 'E': '2', 'I': '3', 'O': '4', 'U': '5'}

// Encode encodes the provided string s.
// Every vowel becomes a corresponding number in the range [1,5].
func Encode(s string) string {
	var builder strings.Builder
	for _, r := range s {
		sub, ok := encodeTable[r]
		if ok {
			builder.WriteRune(sub)
		} else {
			builder.WriteRune(r)
		}
	}
	return builder.String()
}
