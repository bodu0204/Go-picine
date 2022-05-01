package piscine

func IterativePower(nb int, power int) int {
	n := 1
	for ; power > 0; power--{
		n *= nb
	}
	return n
}