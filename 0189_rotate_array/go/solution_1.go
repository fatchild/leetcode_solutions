func rotate(nums []int, k int) {
	// if k is 0 there is not rotation to do
	// if k is equal to the length of the array then the rotation is redundant
	if k == 0 || len(nums) == k {
		return
	}

	// if k is greater than the length of the array then we just need the remainder
	// because otherwide we will be doing redundant operations
	if k > len(nums) {
		k = k % len(nums)
	}

	var offset int = 0
	for i := 0; i < len(nums)-1; i++ {
		var j int = len(nums) - k + i + offset

		// if the swap element index is after the last element
		// then we need to force it back to the starting position
		if j == len(nums)-1 {
			offset -= k
		}

		// if k is 1 then we just swap with the last element on each iteration
		if k == 1 {
			j = len(nums) - k
		}

		fmt.Println("i ", i, ": j ", j)

		// do the swap
		var ie = nums[i]
		nums[i] = nums[j]
		nums[j] = ie

		// if k is half the length of the array and j is now at the last index of the array
		// for the first time, we are done
		if len(nums) == k*2 && j == len(nums)-1 {
			return
		}

		if k%2 == 0 && k > len(nums)/2 && j == len(nums)-1 {
			return
		}
	}
}
