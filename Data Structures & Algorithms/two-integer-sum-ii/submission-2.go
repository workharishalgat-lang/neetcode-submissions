func twoSum(numbers []int, target int) []int {

   left := 0 
   right := len(numbers)-1


   for right > left {
     if numbers[left] +  numbers[right] == target {
		return []int{left +1, right+1}
	 }

	 // that means on the left side we might have 
	 // possiblity of finding number
	 if numbers[left] +  numbers[right] > target {
		right--
	 }
     
	 if numbers[left] +  numbers[right] < target {
		left++
	 }
   }


   return []int{-1,-1}



}
