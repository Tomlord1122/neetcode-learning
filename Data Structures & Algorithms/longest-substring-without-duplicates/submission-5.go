func lengthOfLongestSubstring(s string) int {
    charMap := make(map[byte]bool)
    l := 0
    res := 0
    for r := 0; r < len(s); r++{
        for charMap[s[r]] == true{
            // remove and update left
            charMap[s[l]] = false
            l++
        }
        charMap[s[r]] = true
        res = max(res, r - l + 1)
    }
    return res
}
