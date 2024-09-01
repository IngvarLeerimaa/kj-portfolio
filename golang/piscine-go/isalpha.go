package piscine

func IsAlpha(str string) bool {
	x := []rune(str) // slice string
	for i := 0; i <= StrLen(str)-1; i++ {
		if (x[i] >= 0) && (x[i] <= 47) || (x[i] >= 58) && (x[i] <= 64) || (x[i] >= 91) && (x[i] <= 96) || (x[i] >= 123) && (x[i] <= 127) {
			return false // välistab kõik muud asjad peale numbrite ja tähtede. Ascii kood lihtsalt harjutamise prst
		}
	}
	return true // kui muu ss jrelikult tõene.
}
