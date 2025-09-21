package gostudy

func plusOne(digits []int) []int {
	length := len(digits)
	result := 0
	for i := 0; i < len(digits); i++ {
		result += digits[len(digits)-i-1] * power(i)
		if i == len(digits)-1 {
			if digits[i] == 9 {
				length++
			}
		}
	}

	result++
	arr := make([]int, length)
	for i := 0; i < length; i++ {
		arr[length-1-i] = result % 10
		result = result / 10
	}
	return arr
}

func power(count int) int {
	result := 1
	for i := 0; i < count; i++ {
		result = result * 10
	}
	return result
}
