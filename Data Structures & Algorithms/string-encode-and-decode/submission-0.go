type Solution struct{}

func (s *Solution) Encode(strs []string) string {
    // [length]#[str]
    res := []byte{}
    for _, str := range strs{
        res = append(res, strconv.Itoa(len(str))...)
        res = append(res, '#')
        res = append(res, str...)
    }
    return string(res)
}

func (s *Solution) Decode(encoded string) []string {
    decoded := []string{}

    // Find the length and skip '#'
    // extract str and append to decoded
    i := 0
    for i < len(encoded){
        j := i
        for j < len(encoded) && encoded[j] != '#'{
            j++
        }
        length, _ := strconv.Atoi(encoded[i:j])
        j++ // skip '#'
        // Append str to decoded
        decoded = append(decoded, encoded[j:j+length])
        // update i
        i = j + length
    }
    return decoded
}