package piscine

import (
	"ft"
)

//func PrintStr(s string) {
//		for i := 0; i < len(s); i++{
//			ft.PrintRune((rune)(s[i]))
//	}
//}

func PrintStr(s string) {
	for _, r := range s{
		ft.PrintRune(r)
	}
}