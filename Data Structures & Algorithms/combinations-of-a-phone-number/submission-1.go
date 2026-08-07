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
	var bt func(i int, curStr string)
	bt = func(i int, curStr string){
		if len(curStr) == len(digits){
			res = append(res, curStr)
			return
		}
		for _, d := range digitMap[digits[i]]{
			bt(i+1, curStr+string(d))
		}
	}
	bt(0,"")
	return res
}
