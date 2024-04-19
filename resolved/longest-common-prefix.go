package resolved

import (
	"strings"
)

func LongestCommonPrefix(strs []string) string {
	prefix := ""
	if len(strs) == 0 {
		return prefix
	}

	for idx := range strs[0] {
		searchedPrefix := strs[0][:idx+1]

		wordsListLength := len(strs)
		wordsWithPrefix := 0
		for _, word := range strs {
			if hasPrefix := strings.HasPrefix(word, searchedPrefix); hasPrefix {
				wordsWithPrefix++
			}
		}

		if wordsWithPrefix == wordsListLength {
			prefix = searchedPrefix
		}
	}

	return prefix
}
