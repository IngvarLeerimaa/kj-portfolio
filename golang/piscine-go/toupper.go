package piscine

func ToUpper(s string) string {
	x := []rune(s)                   // slice slice
	result := ""                     // vahed
	for y := 0; y <= len(s)-1; y++ { // k'ib selle vahemiku l'bi
		if (x[y] >= 'a') && (x[y] <= 'z') { // t'hestik
			x[y] = x[y] - 32 // teeb suureks t'heks
		}
		result += string(x[y])
	}
	return result
}
