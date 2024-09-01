package piscine

func IsLower(s string) bool {
	x := []byte(s)        // teeme byteks et saaks üle selle rangedega
	y := 0                // counter
	for _, z := range x { // rangeb üle xi
		if z >= 'a' && z <= 'z' { // kui antud koht on suurem või võrdne ja vaiksem
			y++ // counter +1
		}
	}
	if y == len(s) { // counter == len(s)
		return true
	} else {
		return false
	}
}
