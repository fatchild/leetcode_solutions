func hIndex(citations []int) int {
	var citationCount = make([]int, len(citations)+1)

	// create an array which countes the occurances of the input array element values
	for i := 0; i < len(citations); i++ {
		citationCount[min(citations[i], len(citationCount)-1)]++
	}
	fmt.Println(citationCount)

	// iteratate through
	var papers int = 0
	for i := len(citationCount) - 1; i > 0; i-- {
		papers += citationCount[i] // keep track of how many papers the author has produces as this is what h-value compares

		if papers >= i {
			return i // i is the number of papers with x vitations
		}
	}
	return 0
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
