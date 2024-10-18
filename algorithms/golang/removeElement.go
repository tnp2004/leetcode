package removeElement

func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return len(nums)
	}

	if val == nums[len(nums)-1] {
		nums = nums[:len(nums)-1]
	}

	for i := 0; i < len(nums); i++ {
		if nums[i] == val {
			nums = append(nums[:i], nums[i+1:]...)
			i--
		}
	}

	return len(nums)
}
