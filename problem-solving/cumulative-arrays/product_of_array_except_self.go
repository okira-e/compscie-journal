func productExceptSelf(nums []int) []int {
	prefixArr := make([]int, len(nums))
	postfixArr := make([]int, len(nums))

	for i, val := range nums {
		prevValInPrefix := 1
		if i != 0 {
			prevValInPrefix = prefixArr[i-1]
		}

		prefixArr[i] = val * prevValInPrefix
	}

	for i := len(nums) - 1; i >= 0; i-- {
		prevValInPostfix := 1
		if i != len(nums)-1 {
			prevValInPostfix = postfixArr[i+1]
		}

		postfixArr[i] = nums[i] * prevValInPostfix
	}

	answers := make([]int, len(nums))
	for i, _ := range nums {
		prevInPrefix := 1
		nextInPostfix := 1
		if i != 0 {
			prevInPrefix = prefixArr[i-1]
		}
		if i != len(nums)-1 {
			nextInPostfix = postfixArr[i+1]
		}

		answers[i] = prevInPrefix * nextInPostfix
	}

	return answers
}
