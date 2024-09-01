package piscine

func IsNumeric(str string) bool {
	x := []rune(str) // slice string
	for i := 0; i <= StrLen(str)-1; i++ {
		if (x[i] <= 47) || (x[i] >= 58) {
			return false // välistab kõik muud asjad peale numbrite. Ascii kood lihtsalt harjutamise prst
		}
	}
	return true // kui muu ss jrelikult tõene.
}
