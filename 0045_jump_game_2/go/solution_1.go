func jump(nums []int) int {
	var reach int = 0
	var current_end int = 0
	var jumps int = 0

	for i := 0; i < len(nums)-1; i++ {
		if reach < nums[i]+i {
			reach = nums[i] + i
		}

		if i == current_end {
			current_end = reach
			jumps++

			if current_end >= len(nums)-1 {
				break
			}
		}
	}

	return jumps
}
