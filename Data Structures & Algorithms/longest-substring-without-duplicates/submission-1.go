func lengthOfLongestSubstring(s string) int {
    seen := make(map[byte]bool)
    l := 0 
    res := 0
    for r := 0; r < len(s); r++{
        for seen[s[r]] == true{
            // remove the left character
            seen[s[l]] = false
            l++
        }
        seen[s[r]] = true
        res = max(res, r - l + 1)
    }
    return res
}
