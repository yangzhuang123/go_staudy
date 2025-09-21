package gostudy

func twoSum(nums []int, target int) []int {
	result := make([]int, 2)
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if i == j {
				continue
			}
			temp := nums[i] + nums[j]
			if temp == target {
				result[0] = i
				result[1] = j
				return result
			}
		}
	}

	test := ""
	su
	return result
}
