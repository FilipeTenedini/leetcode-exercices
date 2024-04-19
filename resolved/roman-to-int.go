package resolved

func RomanToInt(s string) int {
	hashtable := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	sum := 0

	for idx, romanValue := range s {
		if idx != len(s)-1 {
			if hashtable[romanValue] < hashtable[rune(s[idx+1])] {
				sum -= hashtable[romanValue]
			} else {
				sum += hashtable[romanValue]
			}
		} else {
			sum += hashtable[romanValue]
		}
	}
	return sum
}
