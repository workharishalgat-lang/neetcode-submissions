func isPalindrome(s string) bool {
reg := regexp.MustCompile("[^a-zA-Z0-9]+")
	s = reg.ReplaceAllString(s, "")



	sl :=  []rune(s)
		left := 0 
	right := len(s)-1


	for right > left {
	
	    
		if strings.ToLower(string(sl[right])) != strings.ToLower(string(sl[left])) {
          return false
		}
		right--
		left++
	}
	return true


}
