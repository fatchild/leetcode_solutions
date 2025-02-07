func hIndex(citations []int) int {
	if len(citations) == 1 {
		if citations[0] == 0 {
			return 0
		}
		return 1
	}

	citations = reverse(countingSort(citations), 0, len(citations)-1)

	for i := 0; i < len(citations); i++ {
		if citations[i] < i+1 {
			return i
		}
	}

	return len(citations)
}

func countingSort(inputArray []int) []int {
	// Declare an auxiliary array countArray[] of size max(inputArray[])+1 and initialize it with 0.
	var countArray = make([]int, max(inputArray)+1)

	// Traverse array inputArray[] and map each element of inputArray[] as an index of countArray[] array, i.e., execute countArray[inputArray[i]]++ for 0 <= i < N.
	for i := 0; i < len(inputArray); i++ {
		countArray[inputArray[i]]++
	}

	// Calculate the prefix sum at every index of array inputArray[].
	for i := 1; i < len(countArray); i++ {
		countArray[i] += countArray[i-1]
	}

	// Create an array outputArray[] of size N.
	var outputArray = make([]int, len(inputArray))

	// Traverse array inputArray[] from end and update outputArray[ countArray[ inputArray[i] ] – 1] = inputArray[i]. Also, update countArray[ inputArray[i] ] = countArray[ inputArray[i] ]- – .
	for i := len(inputArray) - 1; i >= 0; i-- {
		outputArray[countArray[inputArray[i]]-1] = inputArray[i]
		countArray[inputArray[i]]--
	}

	return outputArray
}

func max(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	var max int = nums[0]

	for i := 1; i < len(nums); i++ {
		if nums[i] > max {
			max = nums[i]
		}
	}

	return max
}

func reverse(nums []int, s int, e int) []int {
	for s < e {
		nums[s], nums[e] = nums[e], nums[s]
		s++
		e--
	}
	return nums
}
