package piscine

func RecursivePower(nb int, power int) int {
	if power <= 0{
		return 1
	}
	return RecursivePower(nb, power - 1) * nb
}
