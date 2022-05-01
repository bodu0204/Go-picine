package piscine

func IterativeFactorial(nb int) int {
	n := 1
	for ; nb > 1; nb-- {
		n *= nb
	}
	return n
}