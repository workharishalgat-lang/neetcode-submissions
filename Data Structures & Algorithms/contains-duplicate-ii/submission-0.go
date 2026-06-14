func containsNearbyDuplicate(nums []int, k int) bool {

	for  i, v  := range nums {
		for j:=i+1; j<len(nums);j++ {

			if v == nums[j]&& j-i <= k {
				return true
			}

		}
	}
	return false

}
