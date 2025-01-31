func reverse(nums []int, s int, e int) {
	for s < e {
		nums[s], nums[e] = nums[e], nums[s]
		s++
		e--
	}
}

func rotate(nums []int, k int) {
	if k == 0 {
		return
	}

	var l int = len(nums)
	k = k % l

	reverse(nums, 0, l-1)
	reverse(nums, 0, k-1)
	reverse(nums, k, l-1)
}
