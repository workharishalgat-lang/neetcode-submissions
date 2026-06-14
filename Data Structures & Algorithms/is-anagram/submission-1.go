import "slices"
func isAnagram(s string, t string) bool {
  

  first := []rune(s)
  second := []rune(t)

  slices.Sort(first)
  slices.Sort(second)

  return string(first) == string(second)

    
}
