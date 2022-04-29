package piscine

import(
	"ft"
)

func PrintComb()  {
	var n int
	var s [3]rune
	for n <= 789 {
		s[0] = '0' + (rune)(n /100)
		s[1] = '0' + (rune)((n % 100) /10)
		s[2] = '0' + (rune)(n % 10)
		if s[0] <  s[1] && s[1] < s[2]{
			for i := 0; i < 3; i++ {
				ft.PrintRune(s[i])
			}
			if !(s[0] == '7' && s[1] == '8' && s[2] == '9'){
				ft.PrintRune(',')
				ft.PrintRune(' ')
			}
		}
		n++
	}
	return
}

