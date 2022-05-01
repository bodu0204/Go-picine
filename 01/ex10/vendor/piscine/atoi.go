package piscine

func Atoi(s string) int {
	var i int
	var n int

	sig := 1
	if(s[i] == '+'){
		i++
	}else if s[i] == '-'{
		sig = -1
		i++
	}
	for ; i < len(s) && s[i] >= '0' && s[i] <= '9'; i++ {
		n *= 10
		n += (int)(s[i] - '0')
	}
	if i < len(s){
		return 0;
	}
	return n * sig
}