package reverse_string

func ReverseString(input string) (output string) {
	strR := []rune(input)
	for i, j := 0, len(strR)-1; i < j; i, j = i+1, j-1 {
		strR[i], strR[j] = strR[j], strR[i]
	}
	output = string(strR)
	return
}
