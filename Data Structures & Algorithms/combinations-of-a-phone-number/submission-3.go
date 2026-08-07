func letterCombinations(digits string) []string {
	if len(digits) == 0{
		return []string{}
	}
	res := []string{}
	digitMap := map[byte]string{
		'2':"abc",
		'3':"def",
		'4':"ghi",
		'5':"jkl",
		'6':"mno",
		'7':"pqrs",
		'8':"tuv",
		'9':"wxyz",
	}

	var dfs func(i int, cur string)
	dfs = func(i int, cur string){
		if len(cur) == len(digits){
			res = append(res, cur)
			return
		}
		for _, c := range digitMap[digits[i]]{
			dfs(i+1, cur+string(c))
		}
	}

	dfs(0, "")
	return res
}
