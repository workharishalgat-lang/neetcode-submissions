func twoSum(nums []int, target int) []int {
    visited :=  make(map[int]int)
	visited[nums[0]] = 0 
    for i:=1;i<len(nums);i++ {
		our :=  target - nums[i]
		if v,ok := visited[our]; ok {
			return []int{v,i}
		}else{
			visited[nums[i]] = i
		}
	}
	return []int{0,0}
}
