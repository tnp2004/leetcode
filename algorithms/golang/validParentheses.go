package validParentheses

func isValid(s string) bool {
	if len(s)%2 != 0 {
		return false
	}

	var stack []rune
	hm := map[rune]rune{
		'(': ')',
		'[': ']',
		'{': '}',
	}

	for _, r := range s {
		stackLen := len(stack)
		if _, ok := hm[r]; ok {
			stack = append(stack, r)
		} else {
			if stackLen < 1 {
				return false
			}
			if r == hm[stack[stackLen-1]] {
				stack = append(stack[:stackLen-1])
				continue
			}

			return false
		}
	}

	return len(stack) == 0
}
