func isValid(s string) bool {
    if len(s)%2 != 0 {
		return false
	}

	stack := make([]string,len(s), len(s))
	top := -1

    for _, vv := range s {
		v :=  string(vv)
		if top == -1  {
			if  v == "}" || v == "]" || v == ")" {
return false
			}else{
				top++
				stack[top] = v
			}
			
		}else{
			topE := stack[top]
			if v == "{" || v == "[" || v == "(" {
				top++
				stack[top] = v
			}else{
				if topE == "{" && v == "}"  || topE == "[" && v == "]" || topE == "(" && v == ")" {
					top--
				} else{
					return false
				}
			}
		}
	}
	return top==-1

}
