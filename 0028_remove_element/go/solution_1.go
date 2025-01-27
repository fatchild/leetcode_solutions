func removeElement(nums []int, val int) int {
	var k int = 0
	var l int = 0

	for i := 0; i < len(nums); i++ {
		var e = nums[i]
		if e != val {
			nums[l] = e
			l++
			k++
		}
	}

	return k
}
