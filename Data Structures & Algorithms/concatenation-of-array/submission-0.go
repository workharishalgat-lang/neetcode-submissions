func getConcatenation(nums []int) []int {
    out :=  make([]int, len(nums)*2)

    for i:=0;i< len(nums);i++ {
        out[i] =  nums[i] 
        out[i+len(nums)] =  nums[i]
    }
    return out
    
}
