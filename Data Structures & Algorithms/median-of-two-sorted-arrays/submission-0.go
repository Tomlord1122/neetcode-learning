// findMedianSortedArrays returns the median of two sorted arrays.
// Binary search on the partition of the shorter array.
// Time: O(log(min(m,n))), Space: O(1)
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	A, B := nums1, nums2
	// Ensure A is the shorter array to keep search space small.
	if len(A) > len(B) {
		A, B = B, A
	}

	m, n := len(A), len(B)
	total := m + n
	half := total / 2

	l, r := 0, m // search i in [0..m]
	for {
		i := (l + r) / 2        // partition index in A (count on left side from A)
		j := half - i           // partition index in B (count on left side from B)

		// Boundaries: if i==0, Aleft = -inf; if i==m, Aright = +inf
		var Aleft int
		if i > 0 {
			Aleft = A[i-1]
		} else {
			Aleft = -1 << 60 // effectively -infinity
		}
		var Aright int
		if i < m {
			Aright = A[i]
		} else {
			Aright = 1 << 60 // effectively +infinity
		}

		var Bleft int
		if j > 0 {
			Bleft = B[j-1]
		} else {
			Bleft = -1 << 60
		}
		var Bright int
		if j < n {
			Bright = B[j]
		} else {
			Bright = 1 << 60
		}

		// Check if partitions are valid
		if Aleft <= Bright && Bleft <= Aright {
			// Found correct partition
			if total%2 == 1 {
				// odd length -> min of right side
				if Bright < Aright {
					return float64(Bright)
				}
				return float64(Aright)
			}
			// even length -> average of max(left) and min(right)
			leftMax := Aleft
			if Bleft > leftMax {
				leftMax = Bleft
			}
			rightMin := Aright
			if Bright < rightMin {
				rightMin = Bright
			}
			return float64(leftMax+rightMin) / 2.0
		}

		// Adjust search space
		if Aleft > Bright {
			// A's left too big, move i left
			r = i - 1
		} else {
			// B's left too big, move i right
			l = i + 1
		}
	}
}
