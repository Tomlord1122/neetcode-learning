func partition(s string) [][]string {
	cur := []string{}
	res := [][]string{}

	var dfs func(i int)
	dfs = func(i int){
		if i == len(s){
			res = append(res, append([]string{}, cur...))
			return 
		}

		for j := i; j < len(s); j++{
			if isPalidrome(s[i:j+1]){
				cur = append(cur, s[i:j+1])
				dfs(j+1)
				cur = cur[:len(cur)-1]
			}
		}
	}
	dfs(0)
	return res
}


func isPalidrome(s string) bool{
	i, j := 0, len(s)-1
	for i < j{
		if s[i] != s[j]{
			return false
		}
		i++
		j--
	}
	return true
}