package piscine

func Sqrt(nb int) int {
	r := 1

	for i := 2; i * i <= nb; i++ {
		if nb % i * i == 0{
			nb /= i * i
			r *= i
			i = 2
		}
	}
	if (nb == 1){
		return r
	}
	return 0
}