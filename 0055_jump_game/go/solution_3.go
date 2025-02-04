func canJump(nums []int) bool {
	if len(nums) == 1 {
		return true
	}

	var needs int = 1
	for i := len(nums) - 2; i >= 0; i-- {
		if nums[i] >= needs {
			needs = 1
		} else {
			needs++
		}
	}

	return nums[0] >= needs
}
