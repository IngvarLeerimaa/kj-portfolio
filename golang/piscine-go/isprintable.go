package piscine

func IsPrintable(s string) bool {
	x := []rune(s) // slice string
	for i := 0; i <= StrLen(s)-1; i++ {
		if (x[i] <= 31) || (x[i] >= 127) {
			return false // välistab kõik muud asjad peale numbrite. Ascii kood lihtsalt harjutamise prst
		}
	}
	return true // kui muu ss jrelikult tõene.
}
