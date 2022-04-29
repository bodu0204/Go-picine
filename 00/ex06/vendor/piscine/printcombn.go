package piscine

import(
	"ft"
)

type nums [9]int

func (i *nums) terat(n int) bool {
	ii := n - 1
	for ;ii > 0 && i[ii] > (10 - n) + ii; ii--{
		i[ii - 1]++
	}
	for ; ii + 1 < n; ii++ {
		i[ii + 1] = i[ii] + 1
	}
	if i[n - 1] > 9{
		return false
	}else {
		return true
	}
}

func (i *nums) out(n int) {
	for ii := 0; ii < n; ii++ {
		ft.PrintRune('0' + (rune)(i[ii]))
	}
}

func PrintCombN(n int) {
	var i nums

	if (n < 1 || n > 9){
		return
	}
	for ii := 0; ii < n; ii++ {
		i[ii] = ii
	}
	for ii := 0; i.terat(n); ii++{
		if ii > 0{
			ft.PrintRune(',')
			ft.PrintRune(' ')
		}
		i.out(n)
		i[n - 1]++
	}
}
