type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	encoded := []byte{}
	for _, str := range strs{
		encoded = append(encoded, strconv.Itoa(len(str))...)
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
		decoded = append(decoded, encoded[j:j+length])
		i = j + length
	}

	return decoded
}



// i want to encode a single string to [length]#[string]
// so after we encode the array of strings
// the output will be like [length]#[string][length]#[string]
// we can keep tracking of the length and # to decode it well
