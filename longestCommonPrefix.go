package gostudy

func longestCommonPrefix(strs []string) string {

	prefix := ""
	flag := true
	index := 0
	for flag {
		index++
		for i := 0; i < len(strs); i++ {
			if len(strs[i]) < index {
				flag = false
				break
			}

			tempPrefix := strs[i][:index]
			if i == 0 {
				prefix = tempPrefix
			} else {
				if prefix != tempPrefix {
					prefix = prefix[:len(prefix)-1]
					flag = false
					break
				}
			}
		}
	}
	return prefix

}
