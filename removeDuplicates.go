package gostudy

func removeDuplicates(nums []int) int {
	tempValue := nums[0]
	tempIndex := 0

	for i := 1; i < len(nums); i++ {
		if tempValue != nums[i] {
			tempIndex++
			tempValue = nums[i]
			nums[tempIndex] = tempValue
		}
	}

	return tempIndex + 1
}
