func characterReplacement(s string, k int) int {
    charCount := make(map[byte]int) // (char, freq)
    res, l := 0, 0
    for r := 0; r < len(s); r++{
        charCount[s[r]]++
        for r - l + 1 - getMaxFreq(charCount) > k{
            charCount[s[l]]--
            l++
        }
        res = max(res, r - l + 1)
    }
    return res
}


func getMaxFreq(charCount map[byte]int) int{
    maxF := 0
    for _, freq := range charCount{
        if freq > maxF{
            maxF = freq
        }
    }
    return maxF
}