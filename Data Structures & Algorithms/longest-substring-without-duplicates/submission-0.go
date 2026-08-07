func lengthOfLongestSubstring(s string) int {
    seen := make(map[byte]bool)
    l := 0
    res := 0
    for r := 0; r < len(s); r++{
        for seen[s[r]] == true{
            seen[s[l]] = false
            l++
        }
        seen[s[r]] = true
        res = max(res, r - l + 1)
    }
    return res
}

// Use a hashMap to check the char is seen or not
// The remain part we can use two pointer to solve this question
