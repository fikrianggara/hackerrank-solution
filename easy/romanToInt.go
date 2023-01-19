package main

import "fmt"

func romanToInt(s string) int {
	romanChar := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}
	var result int = 0
	for i := 0; i < len(s); i++ {
		//selama masih ada
		if i < len(s)-1 {
			switch string([]rune(s)[i]) + string([]rune(s)[i+1]) {
			case "IV":
				result = result + 4
				i = i + 1
			case "IX":
				result = result + 9
				i = i + 1
			case "XL":
				result = result + 40
				i = i + 1
			case "XC":
				result = result + 90
				i = i + 1
			case "CD":
				result = result + 400
				i = i + 1
			case "CM":
				result = result + 900
				i = i + 1
			default:
				result = result + romanChar[string([]rune(s)[i])]
			}
		} else {
			result = result + romanChar[string([]rune(s)[i])]
		}
	}
	return result
}

func main() {

	fmt.Println(romanToInt("IV"))
}
