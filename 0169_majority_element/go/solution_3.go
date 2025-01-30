func majorityElement(nums []int) int {
	var n int = len(nums)

	if n == 1 {
		return nums[0]
	}

	var av int = nums[0]
	var ao int = 1

	for i := 1; i < len(nums); i++ {
		if nums[i] == av {
			ao++
		} else {
			ao--
			if ao == 0 {
				av = nums[i]
				ao = 1
			}
		}
	}

	return av
}
