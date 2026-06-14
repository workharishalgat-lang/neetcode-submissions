func hasDuplicate(nums []int) bool {
    temp :=make([]int,0)
   for _,v := range nums {
     if !hasElement(temp,v){
       temp =  append(temp,v)
     }else{
        return true
     }
   }
   return false
}

func hasElement(a []int, key int) bool{
   for _ ,v := range a {
    if v == key {
        return true
    }   
   }
    return false
}
