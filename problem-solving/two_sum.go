func twoSum(nums []int, target int) []int {
	dict := map[int]int{}

	for i, val := range nums {
		if _, ok := dict[target-val]; ok {
			return []int{dict[target-val], i}
		}

		dict[val] = i
	}

	return []int{}

}
