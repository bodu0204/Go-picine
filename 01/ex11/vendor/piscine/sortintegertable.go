package piscine

func SortIntegerTable(table []int) {
	var min int

	for i := 0; i < len(table); i++ {
		min = i
		for ii := i + 1; ii < len(table); ii++ {
			if table[min] > table[ii]{
				min = ii
			}
		}
		table[i], table[min] = table[min], table[i]
	}
}