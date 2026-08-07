func isAnagram(s string, t string) bool {
    matches := [26]int{}
    if len(s) != len(t){
        return false
    }

    for i := 0; i < len(s); i++{
        matches[s[i]-'a']++
    }
    for i := 0; i < len(t); i++{
        matches[t[i]-'a']--
    }

    for i := 0; i < 26; i++{
        if matches[i] != 0{
            return false
        }
    }
    return true
}
