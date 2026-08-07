type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	// encoded those strings into [length]#[string] format
	encoded := []byte{}
	for _, s := range strs{
		length := strconv.Itoa(len(s))
		encoded = append(encoded, length...)
		encoded = append(encoded, '#')
		encoded = append(encoded, s...)
	}
	return string(encoded)
}

func (s *Solution) Decode(encoded string) []string {
	decoded := []string{}

	i := 0
	for i < len(encoded){
		j := i
		for j < len(encoded) && encoded[j] != '#'{
			j++
		}
		// right now j is at the '#' index
		length, _ := strconv.Atoi(encoded[i:j])
		j++
		// get string
		str := encoded[j:j+length]
		decoded = append(decoded, str)
		// update i 
		i = j + length
	}
	return decoded
}
