import "slices"
func groupAnagrams(strs []string) [][]string {

	sorted := make([]string, 0)

	for _, v := range strs {
		 vv := []rune(v)
		 slices.Sort(vv)
		 sorted =  append (sorted, string(vv))

	}

	mapIndex :=  make(map[string][]int)
	// act = 0
    for i , value := range sorted {
		if  _, present :=  mapIndex[value]; present {
           mapIndex[value]  = append(mapIndex[value], i)
		}else{
			mapIndex[value] = []int{i}
		}
	}
    // created the final output
	out :=  [][]string{}
	for _, v :=  range mapIndex {
         d := make([]string,0)
		 for _, index :=  range v {
			d = append(d,strs[index])
		 }
		 out   =  append(out,d)
	}
	return out
    

}
