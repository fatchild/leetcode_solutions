func canJump(nums []int) bool {
	var place int = nums[0]
	var place_idx int = 0
	for place < len(nums) {
		fmt.Println(place, place_idx, nums[place_idx+place])
		if nums[place_idx+place] == 0 || place_idx+place > len(nums)-1 {
			return false
		}
		place += nums[place_idx+place]
	}
	return true
}
