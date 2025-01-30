func majorityElement(nums []int) int {
	var n int = len(nums)

	if n == 1 {
		return nums[0]
	}

	// count
	var hm = map[int]int{}

	for i := 0; i < len(nums); i++ {
		var e int = nums[i]
		val, ok := hm[e]
		if ok {
			hm[e] = val + 1
		} else {
			hm[e] = 1
		}
	}

	// sort
	var ak int = nums[0]
	var av int = 0

	for k, v := range hm {
		if v > av {
			av = v
			ak = k
		}
	}

	return ak
}
