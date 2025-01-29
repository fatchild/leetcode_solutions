// inspired by: db96's answer titled: O(1) Space, look back 2 and no counter - very fast!

func removeDuplicates(nums []int) int {
	var nl = len(nums)

	if nl == 1 {
		return 1
	}

	var k int = 2

	for i := 2; i < nl; i++ {
		var c int = nums[i]
		var km1 int = nums[k-1]
		if nums[i] != km1 || km1 != nums[k-2] {
			nums[k] = c
			k++
		}
	}

	return k
}
