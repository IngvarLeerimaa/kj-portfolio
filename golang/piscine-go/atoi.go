package piscine

func Atoi(s string) int {
	x := 0
	y := []rune(s)
	if len(y) == 0 {
		return 0
	} else if y[0] == '-' {
		x = negAtoi(y)
	} else if y[0] == '+' {
		x = posAtoi(y)
	} else if y[0] >= '0' || y[0] <= '9' {
		x = normalAtoi(y)
	} else {
		x = 0
	}
	return x
}

func negAtoi(s []rune) int {
	x := 0
	for i := 1; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			return 0
		} else {
			x = x*10 + int(s[i]) - '0'
		}
	}
	return -x
}

func posAtoi(s []rune) int {
	x := 0
	for i := 1; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			return 0
		} else {
			x = x*10 + int(s[i]) - '0'
		}
	}
	return x
}

func normalAtoi(s []rune) int {
	x := 0

	for i := 0; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			return 0
		} else {
			x = x*10 + int(s[i]) - '0'
		}
	}
	return x
}
