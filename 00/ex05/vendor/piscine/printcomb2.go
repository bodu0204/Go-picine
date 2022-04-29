package piscine

import(
	"ft"
)

type num int
func (n num) printnum() {
	ft.PrintRune('0' + (rune)(n / 10))
	ft.PrintRune('0' + (rune)(n % 10))
}

func (n num) toptnum() num{
	return n / 100
}
func (n num) btmtnum() num{
	return n % 100
}

func PrintComb2(){
	var i num
	for ; i <= 9899; i++ {
		if i.toptnum() < i.btmtnum(){
			i.toptnum().printnum()
			ft.PrintRune(' ')
			i.btmtnum().printnum()
			if i < 9899{
				ft.PrintRune(',')
				ft.PrintRune(' ')
			}
		}
	}
}
