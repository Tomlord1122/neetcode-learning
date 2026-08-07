func partition(s string) [][]string {
	res := [][]string{}
	cur := []string{}

	var dfs func(i int)
	dfs = func(i int){
		if i >= len(s){
			// append cur to res 
			res = append(res, append([]string{}, cur...))
			return 
		}
		for j := i; j < len(s); j++{
			if isPalindrome(s, i, j){
				cur = append(cur, s[i:j+1])
				dfs(j+1)
				cur = cur[:len(cur)-1]
			}
		}
	}
	dfs(0)
	return res
}

func isPalindrome (s string, l, r int) bool{
	for l < r{
		if s[l] != s[r]{
			return false
		}
		l++
		r--
	}
	return true
}