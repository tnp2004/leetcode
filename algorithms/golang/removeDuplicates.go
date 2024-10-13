package removeDuplicates

func removeDuplicates(s string) string {
	if len(s) < 2 {
		return s
	}

	stack := make([]rune, 0, len(s))
	for _, r := range s {
		stackLen := len(stack)
		if stackLen < 1 || r != stack[stackLen-1] {
			stack = append(stack, r)
			continue
		}
		stack = append(stack[:stackLen-1])
	}

	return string(stack)
}
