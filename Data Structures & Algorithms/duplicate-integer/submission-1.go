func hasDuplicate(nums []int) bool {
    d:= make(map[int]int)
    for _,v:= range nums {
        if _,ok:=d[v]; ok {
            return true
        }else{
            d[v]=v
        }
    }
    return false
}
