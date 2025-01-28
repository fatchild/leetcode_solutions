func bubbleSort(nums []int, t int) []int {
	var sorted bool = true

	for i := 0; i < t-1; i++ {
		var elem int = nums[i]
		var nextElem int = nums[i+1]

		if elem > nextElem {
			sorted = false
			nums[i] = nextElem
			nums[i+1] = elem
		}
	}

	if !sorted {
		bubbleSort(nums, t)
	}

	return nums
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	copy(nums1[m:], nums2)
	nums1 = bubbleSort(nums1, len(nums1))
}
