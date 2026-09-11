func getConcatenation(nums []int) []int {
    out :=  make([]int,0)

    out = append(out, nums...)
    out = append(out, nums...)
    return out
    
}
