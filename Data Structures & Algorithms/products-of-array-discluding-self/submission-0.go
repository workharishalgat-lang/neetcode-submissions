func productExceptSelf(nums []int) []int {
	counter:=0
	out := make([]int, 0)
        for  counter < len(nums) {
			prod:=1
		for j:=0; j<len(nums); j++{
			if j == counter {
				continue
			}else{
              prod = prod*nums[j]
			}
		}
		out =  append(out, prod)
       counter++
		
		
	}
	return out
}
