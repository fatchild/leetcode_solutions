func jump(nums []int) int {
	if len(nums) == 1 {
		return 0
	}

	var target int = len(nums) - 1

	return 1 + JumpUtil(nums, target)
}

func JumpUtil(nums []int, target int) int {
	var max int = 0

	for i := target - 1; i >= 0; i-- {
		if i+nums[i] >= target {
			max = i
		}
	}

	if max == 0 {
		return 0
	}

	target = max

	return 1 + JumpUtil(nums, target)
}
