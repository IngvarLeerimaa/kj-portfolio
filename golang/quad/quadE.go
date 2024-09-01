package piscine

import "github.com/01-edu/z01"

func QuadE(x, y int) {
	// x = b == horisontaal; y = a == vertikaal
	for a := 1; a <= y; a++ {
		// kordab kuni a = y vertikaal
		for b := 1; b <= x; b++ {
			// kordab kuni b == x horisontaal
			if (a == 1 && b == 1) || (a == y && b == x) {
				// 1) vasal ülemine; 2) parem alumine
				z01.PrintRune('A')
				// prindib 'A'
			} else if (a == y && b == 1) || (b == x && a == 1) {
				//vasak alumine, parem ülemine
				z01.PrintRune('C')
				// prindib 'A'
			} else if b == 1 && (a != 1 || a != y) {
				// 1)esimese rea t2rn -  kui on esimene tulp aga mitte 1 v viimane koht, siis prindib 'B'
				z01.PrintRune('B')
			} else if b == x && (a != y && a != 1) {
				// 1)viimase rea B -  kui on viimane tulp aga mitte 1 v viimane koht, siis prindib 'B'
				z01.PrintRune('B')
			} else if a == 1 && (b != x && b != 1) {
				// 	// 1) kui on vasak esimene 1; kuik6ige parem viimane; kui pole parem esimene
				z01.PrintRune('B')
				// 	// prindib 'B'
			} else if a == y && (b != x && b != 1) {
				// 	// 1) kui on parem viimane 1; kui pole parem viimane ; kui pole parem esimene
				z01.PrintRune('B')
			} else {
				// kui ei vasta tingimustele siis prindib tühiku ehk tuleb seest t8hi
				z01.PrintRune(' ')
			}
		}
		z01.PrintRune('\n')
	}
}
