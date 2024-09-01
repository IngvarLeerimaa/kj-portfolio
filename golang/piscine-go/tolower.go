package piscine

func ToLower(s string) string {
	x := []rune(s)                   // slice slice
	result := ""                     // vahed
	for y := 0; y <= len(s)-1; y++ { // k'ib selle vahemiku l'bi
		if (x[y] >= 'A') && (x[y] <= 'Z') { // t'hestik
			x[y] = x[y] + 32 // teeb suureks t'heks
		}
		result += string(x[y])
	}
	return result
}
