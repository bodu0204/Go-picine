package piscine

func StrRev(s string) string {
	var r string

	for _, l := range s{
		r = string(l) + r
	}
	return r
}
