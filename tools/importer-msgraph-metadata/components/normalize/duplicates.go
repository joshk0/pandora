package normalize

import (
	"fmt"
	"regexp"
	"strings"
)

func DeDuplicate(name string) string {
	name2 := regexp.MustCompile(`\b(\w+)\s+$1\b`).ReplaceAllString(name, "$1")
	fmt.Println(name2)
	nameSpaced := regexp.MustCompile("([A-Z])").ReplaceAllString(name, " $1")
	nameParts := strings.Split(strings.TrimSpace(nameSpaced), " ")
	newParts := make([]string, 0, len(nameParts))
	checkPieces := make([]string, 0)
	for i := 0; i < len(nameParts); i++ {
		checkPiece := ""
		for j := 0; j < i; j++ {
			checkPiece = checkPiece + nameParts[j]
		}
		checkPieces = append(checkPieces, checkPiece)
		if i > 0 && strings.EqualFold(nameParts[i], nameParts[i-1]) {
			continue
		}
		newParts = append(newParts, nameParts[i])
	}
	return strings.Join(newParts, "")
}

func DeDuplicateName(name string) string {
	words := make([]string, 0)
	i := 0
	for j, c := range name {
		char := string(c)
		if j > 0 && strings.ToUpper(char) == char {
			i++
		}
		if len(words) <= i {
			words = append(words, "")
		}
		words[i] = words[i] + char
	}

	out := ""
	for k, word := range words {
		if k > 0 && strings.EqualFold(words[k-1], word) {
			continue
		}
		out = out + word
	}

	return out
}
