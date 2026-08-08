func letterCombinations(digits string) []string {
	if len(digits) == 0{
		return []string{}
	}

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

	res := []string{}
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

// Input: digits = "34"
// Output: ["dg","dh","di","eg","eh","ei","fg","fh","fi"]