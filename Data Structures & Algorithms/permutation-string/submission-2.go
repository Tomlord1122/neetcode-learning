func checkInclusion(s1 string, s2 string) bool {
    var s1Map [26]int
    var s2Map [26]int

    for i := 0; i < len(s1); i++{
        s1Map[s1[i] - 'a']++
    }

    if len(s1) > len(s2){
        return false
    }

    for i := 0; i < len(s1); i++{
        s2Map[s2[i]-'a']++
    }


    matches := 0

    for i := 0; i < 26; i++{
        if s1Map[i] == s2Map[i]{
            matches++
        }
    }

    for l, r := 0, len(s1); r < len(s2); l, r = l + 1, r + 1{
        if matches == 26{
            return true
        }
        idx := s2[r] - 'a'
        s2Map[idx]++
        if s1Map[idx] == s2Map[idx]{
            matches++
        } else if s1Map[idx]+1 == s2Map[idx]{
            // It was matched. And right now it wasn't
            matches--
        }
        idx = s2[l] - 'a'
        s2Map[idx]--
        if s1Map[idx] == s2Map[idx]{
            matches++
        } else if s1Map[idx]-1 == s2Map[idx]{
            matches--
        }
    }
    return matches == 26
}



