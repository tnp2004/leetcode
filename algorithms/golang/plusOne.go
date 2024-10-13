package plusOne

func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		v := digits[i]
		if v == 9 {
			digits[i] = 0
		} else {
			digits[i]++
			break
		}

		if digits[0] == 0 {
			digits := append([]int{1}, digits...)
			return digits
		}
	}

	return digits
}
