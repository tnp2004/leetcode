package firstMissingPositive

func firstMissingPositive(nums []int) int {
	var min, max int
	hm := make(map[int]bool, len(nums))
	for _, v := range nums {
		hm[v] = true
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}

	if min > 1 || max < 1 {
		return 1
	}

	if min < 1 {
		min = 1
	}

	for n := min; n < max; n++ {
		if _, ok := hm[n]; !ok {
			return n
		}
	}

	return max + 1
}
