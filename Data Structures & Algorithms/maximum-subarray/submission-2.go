func maxSubArray(nums []int) int {

	largestSum := nums[0]

    for  i:=0;i<len(nums);i++ {
		currentSum:= nums[i]
		for j:=i+1;j<len(nums);j++{
			currentSum +=nums[j]
		  if currentSum > largestSum {
			largestSum = currentSum
		}

		}
		if currentSum > largestSum {
			largestSum = currentSum
		}
 
	

	}
	return largestSum
    


    
}
