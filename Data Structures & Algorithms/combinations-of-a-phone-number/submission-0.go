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
	var backtrack func(i int, curStr string)
	backtrack = func(i int, curStr string){
		if len(curStr) == len(digits){
			res = append(res, curStr)
			return 
		}
		for _, c := range digitMap[digits[i]]{
			backtrack(i+1, curStr+string(c))
		}
	}
	backtrack(0, "")
	return res
}
