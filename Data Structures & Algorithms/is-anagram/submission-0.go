func isAnagram(s string, t string) bool {
    if len(s) != len(t){
        return false
    }
    charCount := [26]int{}
    for i := 0; i < len(s); i++{
        charCount[s[i]-'a']++
    }

    for i := 0; i < len(t); i++{
        charCount[t[i]-'a']--
        if charCount[t[i]-'a'] < 0{
            return false
        }
    }
    return true
}
