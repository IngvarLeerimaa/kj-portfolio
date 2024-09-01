package piscine

func Split(s, sep string) []string {
	var sRunes []rune
	for index := 0; index < len(s); index++ {
		if s[index] == sep[0] && len(sep) <= len(s[index:]) { // if letter of index  == first letter of sep and there is enouh letters(word must be smaller than len)
			for i := 0; i < len(sep); i++ {
				if s[index+i] == sep[i] {
					if i == len(sep)-1 { // has made it to last letter of sep
						index += len(sep) - 1 // skips letters in s
						sRunes = append(sRunes, ' ')
						break
					}
				} else {
					sRunes = append(sRunes, rune(s[index]))
					break
				}
			}
		} else {
			sRunes = append(sRunes, rune(s[index]))
		}
	}
	return SplitWhiteSpaces(string(sRunes))
}
