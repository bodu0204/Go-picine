package piscine

func BasicAtoi2(s string) int {
	var n int

	for _, i := range s{
		if (i < '0' || i > '9'){
			return 0
		}
		n *= 10
		n += (int)(i - '0')
	}
	return n
}