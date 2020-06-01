package numcode

import "strings"

var decodeTable = map[rune]rune{'1': 'a', '2': 'e', '3': 'i', '4': 'o', '5': 'u'}

// Decode decodes the provided string s.
// Every number in the range [1,5] becomes a corresponding lowercase vowel.
func Decode(s string) string {
	var builder strings.Builder
	for _, r := range s {
		sub, ok := decodeTable[r]
		if ok {
			builder.WriteRune(sub)
		} else {
			builder.WriteRune(r)
		}
	}
	return builder.String()
}
