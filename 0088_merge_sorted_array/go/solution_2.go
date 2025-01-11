func merge(nums1 []int, m int, nums2 []int, n int) {
	// Protect in the instances of either of the arrays being empty
	if n == 0 {
		return
	}
	if m == 0 {
		copy(nums1, nums2)
		return
	}

	// no need to make more array accesses then needed
	var a int = 0
	var b int = 0
	var aLast bool = false
	var bLast bool = false
	var nums []int

	for i := 0; i < m+n; i++ {
		var numa int
		if !aLast {
			numa = nums1[a]
		}
		var numb int
		if !bLast {
			numb = nums2[b]
		}

		if aLast {
			nums = append(nums, numb)
			b++
		} else if bLast {
			nums = append(nums, numa)
			a++
		}

		if !aLast && !bLast {
			if numa >= numb {
				nums = append(nums, numb)
				b++
				if b == n {
					bLast = true
				}
			} else {
				nums = append(nums, numa)
				a++
				if a == m {
					aLast = true
				}
			}
		}
	}

	// the newly constructed array is replacing the orignal nums1 array
	copy(nums1, nums)
}