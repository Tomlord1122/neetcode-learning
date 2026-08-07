type Solution struct{}

func (s *Solution) Encode(strs []string) string {
    // length[#]str
    encoded := []byte{}
    for _, s := range strs{
        encoded = append(encoded, strconv.Itoa(len(s))...)
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
        length, _ := strconv.Atoi(encoded[i:j])
        j++
        decoded = append(decoded, encoded[j:j+length])
        i = j + length
    }
    return decoded
}
