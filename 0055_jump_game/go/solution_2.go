func canJump(nums []int) bool {
	if len(nums) == 1 {
		return true
	}

	var pos int = len(nums) - 1
	var path bool

	for i := pos - 1; i >= 0; i-- {
		if nums[i]+i >= pos {
			path = true
			pos = i
		} else {
			path = false
		}
	}

	return path
}
