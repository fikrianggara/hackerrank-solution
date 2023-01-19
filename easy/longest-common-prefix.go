package main

import "strings"

func longestCommonPrefix(strs []string) string {
	currMinLenStr := 0
	minStr := ""
	commonPref := ""

	//1. cari string terpendek,
	//2. kalau ngga ada kombinasi karakter di string terpendek-
	//      yang memiliki common prefix dengan string lain, maka tidak ada common prefix

	for i := 0; i < len(strs); i++ {
		if len(strs[i]) < currMinLenStr {
			currMinLenStr = len(strs[i])
			minStr = strs[i]
		}
	}
	//kalo string terpendek berupa string kosong, tidak ada common prefix
	if currMinLenStr == 0 {
		return ""
	}

	thereIsCommonPref := false
	arrCommonPref := strings.Split(commonPref, "")

	if currMinLenStr == 1 {

		for _, item := range strs {
			if strings.Contains(item, minStr) {
				thereIsCommonPref = true
			} else {
				thereIsCommonPref = false
			}
		}

		if thereIsCommonPref {
			return minStr
		}
	}

	for i := 0; i < len(arrCommonPref)-1; i++ {
		commonPref = commonPref + string([]rune(arrCommonPref[i]))
		for j := 1; j < len(arrCommonPref); j++ {
			commonPref = commonPref + string([]rune(arrCommonPref[j]))

			for _, item := range strs {
				if strings.Contains(item, commonPref) {
					thereIsCommonPref = true
				} else {
					thereIsCommonPref = false
					break
				}
			}

			if !thereIsCommonPref {
				commonPref = ""
				break
			}

		}
	}

	return minStr
}
