package resolved

func TwoSum(numbers []int, target int) []int {
	hashTable := make(map[int]int)

	for i, v := range numbers {
		dif := target - v

		if idx, ok := hashTable[dif]; ok {
			return []int{idx, i}
		}
		hashTable[v] = i
	}
	return nil
}
