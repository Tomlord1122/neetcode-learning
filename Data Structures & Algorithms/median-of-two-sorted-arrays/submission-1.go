func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    a, b := nums1, nums2
    
    // we only want to a be the smaller array
    if len(a) > len(b){
        a, b = b, a
    }

    m, n := len(a), len(b)
    total := m + n
    half := total / 2
    // Example: half =  (6 + 5) / 2 = 5
    // We want the six element
    l, r := 0, m
    for l <= r{
        i := (l + r) / 2
        j := half - i 
        Aleft := math.MinInt
        if i > 0{
            Aleft = a[i-1]
        }
        Aright := math.MaxInt
        if i < m{
            Aright = a[i]
        }
        Bleft := math.MinInt
        if j > 0{
            Bleft = b[j-1]
        }
        Bright := math.MaxInt
        if j < n{
            Bright = b[j]
        }

        if Aleft <= Bright && Bleft <= Aright{
            // Found correct partition
            if total % 2 == 1{
                return min(float64(Bright), float64(Aright))
            }
            leftMax := max(float64(Aleft), float64(Bleft))
            rightMin := min(float64(Aright), float64(Bright))
            return (leftMax + rightMin) /2
        }
        if Aleft > Bright{
            r = i - 1
        } else {
            l = i + 1
        }

    }

    return -1
}   
