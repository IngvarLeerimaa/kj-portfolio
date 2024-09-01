package piscine

func Index(s string, toFind string) int {
	a := []rune(s)              // kust otsime
	b := []rune(toFind)         // mida otsime
	c := len(a)                 // s stringi pikkus
	d := len(b)                 // toFind stringi pikkus
	for i := 0; i <= c-d; i++ { // counter
		if toFind == s[i:i+d] { // v6rdleb slice index kuni vääring
			return (i)
		}
	}
	return -1
}
