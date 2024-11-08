package iteration

func Repeat(char string, n int) string {
	if char == "" {
		return ""
	}

	var str string
	for i := 1; i <= n; i++ {
		str += char
	}

	return str
}
