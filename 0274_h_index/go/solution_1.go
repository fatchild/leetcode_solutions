func hIndex(citations []int) int {
	if len(citations) == 1 && citations[0] == 0 {
		return 0
	}

	citations = bubbleSort(citations)

	for i := 0; i < len(citations); i++ {
		if citations[i] < i+1 {
			return i
		}
	}
	return len(citations)
}

func bubbleSort(nums []int) []int {
	var sorted bool = true

	for i := 0; i < len(nums)-1; i++ {
		var elem int = nums[i]
		var nextElem int = nums[i+1]

		if elem < nextElem {
			sorted = false
			nums[i] = nextElem
			nums[i+1] = elem
		}
	}

	if !sorted {
		bubbleSort(nums)
	}

	return nums
}
