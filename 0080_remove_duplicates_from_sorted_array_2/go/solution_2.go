func removeDuplicates(nums []int) int {
	var k int = 1
	var nl = len(nums)

	if nl == 1 {
		return k
	}

	var s int = 0 // occurances of the last unique element
	var l int = nums[0]

	if l == nums[1] {
		s++
	}

	for i := 1; i < nl; i++ {
		var n int = nums[i]

		if n != l {
			nums[k] = n
			k++
			s = 1
			l = n
		} else if s < 2 {
			nums[k] = n
			k++
			s++
		}
	}

	return k
}
