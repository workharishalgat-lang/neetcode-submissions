func twoSum(nums []int, target int) []int {
	data := map[int]int{}
    
   data[nums[0]] = 0 
   for i:=1;i<len(nums);i++ {
	  v := target - nums[i]
	  if j, present := data[v];  present {
		return []int{j,i}
	  }else{
		data[nums[i]] = i
	  } 
   }
   return []int{0,0}
    
}
