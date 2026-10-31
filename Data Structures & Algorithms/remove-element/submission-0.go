func removeElement(nums []int, val int) int {

i :=0
j :=1

if len(nums) ==0 {
	return 0
}

for j< len(nums) {
	if (nums[j]==nums[i] && nums[i]==  val){
		j++
	}else if(nums[i] == val && nums[j] != val){
		nums[i], nums[j] = nums[j],nums[i]
		i++
		j = i+1
	}else{
		i++
		j = i+1
	}
}
count:=0
for i:=0;i<len(nums);i++{
	if nums[i] == val {
		break
	}
	count++
}
return count

}

