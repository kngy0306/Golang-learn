package myculc

func plus(n1, n2 int) int {
	return n1 + n2
}

func substract(n1, n2 int) int {
	if n1 > n2 {
		return n1 - n2
	} else {
		return n2 - n1
	}
}
