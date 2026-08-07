func addBinary(a string, b string) string {
	res := []byte{}
	car := 0

	aBytes := []byte(a)
	bBytes := []byte(b)

	// reverse it at first
	for i, j := 0, len(aBytes) - 1; i < j; i, j = i+1, j-1{
		aBytes[i], aBytes[j] = aBytes[j], aBytes[i]
	}
	for i, j := 0, len(bBytes) - 1; i < j; i, j = i+1, j-1{
		bBytes[i], bBytes[j] = bBytes[j], bBytes[i]
	}

	n := max(len(aBytes), len(bBytes))

	for i := 0; i < n; i++{
		digitA := 0
		digitB := 0
		if i < len(aBytes){
			digitA = int(aBytes[i] - '0')
		}
		if i < len(bBytes){
			digitB = int(bBytes[i] - '0')
		}

		total := digitA + digitB + car
		res = append(res, byte(total%2) + '0')
		car = total / 2
	}

	if car > 0{
		res = append(res, '1')
	}

	// reverse the res back to get the correct order
	for i, j := 0, len(res) - 1; i < j; i, j = i+1, j-1{
		res[i], res[j] = res[j], res[i]
	}
	return string(res)
}
