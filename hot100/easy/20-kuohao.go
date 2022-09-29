package easy

// isValid 有效的括号，正确的括号
func isValid(s string) bool {
	length := len(s)
	if length%2 != 0 {
		return false
	}
	mapKey := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}
	stack := make([]byte, 0)
	for i := 0; i < length; i++ {
		if s[i] == '{' || s[i] == '(' || s[i] == '[' {
			stack = append(stack, s[i])
		} else if len(stack) > 0 && stack[len(stack)-1] == mapKey[s[i]] {
			stack = stack[:len(stack)-1]
		} else {
			return false
		}
	}

	return len(stack) == 0
}
func test(s string) bool {
	length := len(s)
	if length%2 == 1 {
		return false
	}

	mapKey := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}
	stack := make([]byte, 0)
	for i := 0; i < length; i++ {
		if s[i] == '{' || s[i] == '[' || s[i] == '(' {
			stack = append(stack, s[i])
		} else if len(stack) > 0 && mapKey[s[i]] == stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		} else {
			return false
		}
	}

	return len(stack) == 0
}
