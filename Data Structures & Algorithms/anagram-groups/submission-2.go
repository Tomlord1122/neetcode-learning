func groupAnagrams(strs []string) [][]string {
    res := [][]string{}

    anagrams := make(map[[26]int][]string)
    for _, str := range strs{
        var charCount [26]int
        for i := 0; i < len(str); i++{
            charCount[str[i]-'a']++
        }
        // Append str to anagrams based on charCount
        anagrams[charCount] = append(anagrams[charCount], str)
    }

    for _, anagram := range anagrams{
        res = append(res, anagram)
    }
    return res
}
