package findTheIndexOfTheOccurrenceInAString

func strStr(haystack string, needle string) int {
	n := len(needle)
	h := len(haystack)
	if n > h {
		return -1
	}

	if n == 1 && h == 1 {
		return 0
	}

	for i := 0; i < h; i++ {
		if haystack[i] == needle[0] {
			for j := 0; j < h; j++ {
				if j+i >= h {
					return -1
				}
				if haystack[j+i] != needle[j] {
					break
				}

				if j+1 >= n {
					return i
				}

			}
		}

	}

	return -1
}
