func groupAnagrams(strs []string) [][]string {
    res := [][]string{}
    groups := make(map[[26]int][]string)
    for _, str := range strs{
        var countMap [26]int
        for i := 0; i < len(str); i++{
            countMap[str[i]-'a']++
        }
        groups[countMap] = append(groups[countMap], str)
    }

    // Right now the groups contain all the different anagram group
    // We should append them into res
    for _, group := range groups{
        res = append(res, group)
    }
    return res
}
