func gdc(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func rotate(nums []int, k int) {
	var l int = len(nums)

	k = k % l
	k = l - k // we are left rotating which is equal to right rotation of k = k - n

	var cycles int = gdc(l, k)

	for i := 0; i < cycles; i++ {
		var start_e int = nums[i]
		var cur_i int = i

		for true {
			var next_i int = (cur_i + k) % l

			if next_i == i {
				break
			}

			nums[cur_i] = nums[next_i]
			cur_i = next_i
		}
		nums[cur_i] = start_e
	}
}
