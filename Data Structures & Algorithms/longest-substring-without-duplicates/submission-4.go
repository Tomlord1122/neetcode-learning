func lengthOfLongestSubstring(s string) int {
    seen := make(map[byte]bool)
    res := 0
    l := 0
    for r := 0; r < len(s); r++{
        for seen[s[r]]{
            seen[s[l]] = false
            l++
        }
        seen[s[r]] = true
        res = max(res, (r-l+1))
    }
    return res
}
