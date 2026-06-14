func findMaxConsecutiveOnes(nums []int) int {

	max:=0
	count:=0
	for _,v := range nums {

		if v == 1 {
			count++
			
		}else{
			count =0
		}
        if count > max {
			max = count
		
		}
		

	}
	return max
	
}
