package resolved

func ValidParentheses(s string) bool {
	validBrackets := map[string]string{
		")": "(",
		"}": "{",
		"]": "[",
	}

	queue := []string{}

	for i, bracket := range s {
		if _, ok := validBrackets[string(s[i])]; !ok {
			queue = append(queue, string(bracket))
		} else {
			if len(queue) == 0 {
				return false
			}

			lastBracket := queue[len(queue)-1]
			if lastBracket != validBrackets[string(bracket)] {
				return false
			} else {
				queue = queue[:len(queue)-1]
			}
		}
	}

	return len(queue) <= 0
}
