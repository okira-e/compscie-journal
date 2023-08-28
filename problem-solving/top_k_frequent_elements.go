package main

func topKFrequent(nums []int, k int) []int {
	// -- Populate a map with every number and its occurrence.
	frequencyMap := map[int]int{}
	for _, num := range nums {
		if frequency, ok := frequencyMap[num]; ok {
			frequencyMap[num] = frequency + 1
		} else {
			frequencyMap[num] = 1
		}
	}

	output := make([]int, k)
	newNums := make([]int, len(nums))
	copy(newNums, nums)
	for i := 0; i < k; i += 1 {
		mostFrequentNum := getTopFrequent(frequencyMap)
		// -- Add the KthTopFrequent number to the output.
		output[i] = mostFrequentNum
		// -- Delete the mostFrequentNum from frequencyMap
		delete(frequencyMap, mostFrequentNum)
	}

	return output
}

func getTopFrequent(mp map[int]int) int {
	max := -1
	var mostFrequentNum int
	for num, frequency := range mp {
		if frequency > max {
			max = frequency
			mostFrequentNum = num
		}
	}

	return mostFrequentNum
}
