func removeDuplicates(nums []int) int {
	var k int = 1
	var l int = 1
	var s int = nums[0]

	for i := 1; i < len(nums); i++ {
		var e = nums[i]
		if e != s {
			nums[l] = e
			l++
			k++
		}
		s = e
	}

	return k
}
