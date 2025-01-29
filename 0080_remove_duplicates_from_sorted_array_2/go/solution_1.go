func removeDuplicates(nums []int) int {
	var k int = 1
	var s int = 0 // occurances of the last unique element

	if len(nums) == 1 {
		return k
	}

	if nums[0] == nums[1] {
		s++
	}

	for i := 1; i < len(nums); i++ {
		var l int = nums[i-1]
		var n int = nums[i]

		if n != l {
			nums[k] = n
			k++
			s = 1
		} else if s < 2 {
			nums[k] = n
			k++
			s++
		}
	}

	return k
}
