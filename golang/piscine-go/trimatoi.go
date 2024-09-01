package piscine

/*Write a function that transforms numbers within a string, into an int.

If the - sign is encountered before any number it should determine the sign of the returned int.

This function should only return an int. In the case of an invalid input, the function should return 0.

Note: There will never be more than one sign in a string in the tests.
*/
func TrimAtoi(s string) int {
	i := 0                // counter
	for _, x := range s { // 1st range loop goes through x indexes 1 by 1 and looks for the numbers

		if x >= '0' && x <= '9' {
			if i == '0' { // goes through all possible values from '0' to '9'
				i += 0  // adds the value
				i *= 10 // makes possible to add next didget cuz 1+2 == 3 but we need 1 + 1 = 12
			} else if x == '1' {
				i += 1
				i *= 10
			} else if x == '2' {
				i += 2
				i *= 10
			} else if x == '3' {
				i += 3
				i *= 10
			} else if x == '4' {
				i += 4
				i *= 10
			} else if x == '5' {
				i += 5
				i *= 10
			} else if x == '6' {
				i += 6
				i *= 10
			} else if x == '7' {
				i += 7
				i *= 10
			} else if x == '8' {
				i += 8
				i *= 10
			} else if x == '9' {
				i += 9
				i *= 10
			}
		}
	} // range loop is closed
	num := i / 10             // takes number and / by 10 to get the real one because we *10
	mind := 100               // absurdly larger number or the system breaks
	lin := 0                  // counter
	for index, x := range s { // to get minus index
		if x == '-' { // if first finds '-' saves that to  mind index placement number to mind
			mind = index // returns to loop
		}
		if x == '1' || x == '2' || x == '3' || x == '4' || x == '5' || x == '6' || x == '7' || x == '8' || x == '9' {
			lin = index       // saves index to lin saves index number
			if index > mind { // if minus is before the first number print neg number
				return -num
			} else if lin > 0 { // if minus is after
				return num
			}
		}
	}
	return num
}
