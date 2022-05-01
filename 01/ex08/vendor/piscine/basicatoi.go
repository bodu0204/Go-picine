package piscine

func BasicAtoi(s string) int {
	var n int

	for _, i := range s{
		n *= 10
		n += (int)(i - '0')
	}
	return n
}