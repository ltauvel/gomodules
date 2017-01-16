package crypto

import (
	"strings"
)

func prePend(text string, arr []byte) []byte {
	var result = []byte(text)
	for _, i := range arr {
		result = append(result, i)
	}
	return result
}

func formatKeystring(keystring string) []byte {
	keystring = keystring + strings.Repeat("#", 32)
	return []byte(keystring[:32])
}
