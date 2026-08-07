func characterReplacement(s string, k int) int {
    charCount := make(map[byte]int)
    l := 0
    res := 0
    for r := 0; r < len(s); r++{
        charCount[s[r]]++
        for r - l + 1 - topFreq(charCount) > k{
            charCount[s[l]]--
            l++
        }
        res = max(res, r-l+1)
    }
    return res
}

func topFreq(charCount map[byte]int) int{
    res := 0
    for _, freq := range charCount{
        res = max(res, freq)
    }
    return res
}
// (r - l + 1) - (topFreq in that window) <= k