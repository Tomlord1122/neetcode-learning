func lengthOfLongestSubstring(s string) int {
    // sliding window 
    res := 0
    l := 0
    seen := make(map[byte]bool)
    for r := 0; r < len(s); r++{
        for seen[s[r]]{
            // remove and update l
            seen[s[l]] = false
            l++
        }
        seen[s[r]] = true
        res = max(res, r - l +1)
    }
    return res
}
