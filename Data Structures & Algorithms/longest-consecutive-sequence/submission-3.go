func longestConsecutive(nums []int) int {
    // Use a hash set to allow O(1) existence checks for numbers.
    seen := make(map[int]bool)
    res := 0

    // Add all numbers to the set.
    for _, num := range nums {
        seen[num] = true
    }

    // For each number, try to treat it as the start of a consecutive sequence.
    for _, num := range nums {
        // If the previous number exists, this is not the start of a sequence.
        // We only expand sequences from their smallest element to avoid duplicates.
        if seen[num-1] {
            continue
        }

        // num is the start of a sequence; expand forward.
        length := 1
        next := num + 1

        // Count how many consecutive numbers exist after num.
        for seen[next] {
            length++
            next++
        }

        // Update the global maximum sequence length.
        if length > res {
            res = length
        }
    }

    return res
}