
func longestCommonPrefix(strs []string) string {

	ch := ""
   for i:=0;i<len(strs[0]);i++{
	  ch =  ch + string(strs[0][i])
	  for j:=1;j<len(strs);j++ {
		if !strings.Contains(strs[j],ch){
			if len(ch) > 1 {
				return ch[:len(ch)-1]
			}else{
				return ""
			}
		} 
	  }
   }
   return ch

    
    
}
