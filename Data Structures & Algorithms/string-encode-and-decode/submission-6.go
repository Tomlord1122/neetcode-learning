type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	encoded := []byte{}
	// encode each string as length#string
	for _, str := range strs{
		length := strconv.Itoa(len(str))
		encoded = append(encoded, length...)
		encoded = append(encoded, '#')
		encoded = append(encoded, str...)
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
		length, _ := strconv.Atoi(encoded[i:j])
		j++
		// append the string into decoded
		decoded = append(decoded, encoded[j:j+length])
		i = j + length
	}
	return decoded
}
