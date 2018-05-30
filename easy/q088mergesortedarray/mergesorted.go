package mergesorted

func merge(nums1 []int, m int, nums2 []int, n int) {
	// make space at front of nums1
	for i := m - 1; i >= 0; i-- {
		nums1[i+n] = nums1[i]
	}

	i1, i2 := n, 0

	for p1 := 0; p1 < (n + m); p1++ {
		nums1[p1], i1, i2, _ = choose(nums1, i1, nums2, i2)

	}

}

// choose decides which of the requested elements from nums1 and nums2 is larger
// providing the requested indices lie within the slices, and returns the updated indeces
// if they fall outside ok->false
func choose(nums1 []int, i1 int, nums2 []int, i2 int) (max int, ni1, ni2 int, ok bool) {
	ok1 := i1 < len(nums1)
	ok2 := i2 < len(nums2)

	if !ok1 && !ok2 {
		return -1, i1, i2, false
	}

	if !ok1 {
		return nums2[i2], i1, i2 + 1, true
	}

	if !ok2 {
		return nums1[i1], i1 + 1, i2, true
	}

	if nums1[i1] < nums2[i2] {
		return nums1[i1], i1 + 1, i2, true
	}

	return nums2[i2], i1, i2 + 1, true

}
