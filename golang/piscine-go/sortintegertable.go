package piscine

/*Instructions
Write a function that reorders a slice of int in ascending order.
Expected function
func SortIntegerTable(table []int) {

}*/
func SortIntegerTable(table []int) {
	for j := 0; j < len(table); j++ {
		for i := 1; i < len(table); i++ {
			if table[i-1] > table[i] {
				x := table[i]
				table[i] = table[i-1]
				table[i-1] = x
			}
		}
	}
}
