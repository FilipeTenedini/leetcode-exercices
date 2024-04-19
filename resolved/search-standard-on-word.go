package resolved

func SearchStandardOnWord(word string, errorsQtt int) bool {
	const MINIMUN_LENGTH = 1

	if word == "" || len(word) == MINIMUN_LENGTH {
		return false
	}

	hashtable := make(map[string]int)

	for _, character := range word {
		char := string(character)

		/*****
		if _, ok := hashtable[char]; ok {
			hashtable[char]++
		} else {
			hashtable[char] = 1
		}
		is similar to \/
		**/

		hashtable[char]++
	}

	equalCharsQt := hashtable[string(word[0])]

	errors := 0

	for _, val := range hashtable {
		if val != equalCharsQt {
			errors++
			if errors > errorsQtt {
				return false
			}
		}
	}
	return true
}
