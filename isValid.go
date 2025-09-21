package gostudy

func isValid(s string) bool {

	keyValueMap := make(map[byte]byte, 3)
	keyValueMap[')'] = '('
	keyValueMap[']'] = '['
	keyValueMap['}'] = '{'

	stack := []byte{}
	for i := 0; i < len(s); i++ {
		tempValue := keyValueMap[s[i]]
		if tempValue > 0 {
			if len(stack) == 0 || stack[len(stack)-1] != tempValue {
				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, s[i])
		}
	}
	return len(stack) == 0
}
