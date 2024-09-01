package piscine

// import "fmt"

func StrRev(s string) string {
	muutja := []rune(s)
	//	stringiPikkus := 0
	//	for i := range muutja {
	//		i++
	//		stringiPikkus++
	//	} using len(s) instead of the previous shit
	//fmt.Println(s)
	//fmt.Println(muutja)
	//fmt.Println(len(s))
	//for i, v := range muutja {
	//	fmt.Printf("See on index: %v \t see on v22rtus: %v \n", i, v)
	//}
	var tagurpidi string
	for i := len(s) - 1; i >= 0; i-- {
		tagurpidi += string(muutja[i])
	}
	return tagurpidi
}
