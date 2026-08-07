func characterReplacement(s string, k int) int {
    charCount := make(map[byte]int)
    res := 0
    l := 0
    for r := 0; r < len(s); r++{
        // Valid
        charCount[s[r]]++
        for r - l + 1 - topFreq(charCount) > k{
            charCount[s[l]]--
            l++
        }
        res = max(res, r - l + 1)
    }
    return res
}


func topFreq (charCount map[byte]int) int{
    res := 0
    for _, val := range charCount{
        res = max(res, val)
    }
    return res
}


// 1. maintain a window that it can take at most k replacements -> still valid
// 2. when it's not valid. we should shrink the window

// Valid condition -> (r - l + 1) - mostFreqChar in the current window
