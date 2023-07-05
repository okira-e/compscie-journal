func containsDuplicate(nums []int) bool {
	var myMap = make(map[int]bool)

	for _, num := range nums {
		if myMap[num] {
			return true
		}
		myMap[num] = true
	}

	return false
}
