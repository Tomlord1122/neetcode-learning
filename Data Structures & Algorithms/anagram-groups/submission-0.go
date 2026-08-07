func groupAnagrams(strs []string) [][]string {
    anagrams := make(map[[26]int] []string)

    for _, str := range strs{
        charCount := [26]int{}
        for i := 0; i < len(str); i++{
            charCount[str[i]-'a']++
        }
        // Add this string to the anagrams with the specific key
        anagrams[charCount] = append(anagrams[charCount], str)
    }

    // Now we can gather all sublists and append it to a 2-d array (res)
    res := [][]string{}
    for _, val := range anagrams{
        res = append(res, val)
    }
    return res
}




// make a map. The key is [26]int and the value is a []string
// return a [][]string

