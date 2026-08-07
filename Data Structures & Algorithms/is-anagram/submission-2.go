func isAnagram(s string, t string) bool {
    var match [26]int
    if len(s) != len(t){
        return false
    }
    for i := 0; i < len(s); i++{
        match[s[i]-'a']++
        match[t[i]-'a']--
    }

    for i := 0; i < len(match); i++{
        if match[i] != 0{
            return false
        }
    }
    return true
}
