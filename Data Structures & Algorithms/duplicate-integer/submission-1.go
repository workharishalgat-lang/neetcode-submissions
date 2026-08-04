func hasDuplicate(nums []int) bool {
   data := make(map[int]int)

   for _, v :=  range  nums {
	if _, present := data[v];   present {
		return true
	}else{
		data[v] =  1
	}
   } 
   return false
}
