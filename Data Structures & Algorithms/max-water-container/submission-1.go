func maxArea(heights []int) int {

	cMax := 0
     // brute force
	// for i:=0;i<len(heights);i++ {
	// 	for  j:=i+1;j<len(heights);j++ {
	// 		min:=heights[j]
	// 		if heights[j] > heights[i] {
	// 			min = heights[i]
	// 		}
	// 		o := min * (j-i)
	// 		if o > cMax {
	// 			cMax = o
	// 		}
	// 	}
	// } 
	// return cMax


	low:=0
	high:= len(heights)-1

	for high> low {
      
	  minHeight := heights[low]
	  if heights[high] < heights[low] {
		 minHeight =  heights[high]
	  }
	  temp := minHeight * (high-low)
	  if temp > cMax {
		cMax  = temp
	  }
	  if heights[low] < heights[high] {
		low++
	  }else{
		high--
	  }

	}
	return cMax

}
