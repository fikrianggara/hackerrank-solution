package main

import (
	"fmt"
	"strings"
)

func isValid(b string) bool {
	result := false
	prths := map[string]string{
		")": "(",
		"]": "[",
		"}": "{",
		"(": ")",
		"[": "]",
		"{": "}",
	}

	charStack := make([]string, 0)

	for i := 0; i < len(b); i++ {
		//harus genap
		if len(b)%2 != 0 {
			result = false
			break
		}
		//kalo kurung buka tidak memiliki kurung tutupnya dan sebaliknya, maka salah
		if !strings.Contains(b, prths[string([]rune(b)[i])]) {
			result = false
			break
		}

		//kalo char pertama itu kurung tutup atau char terakhir kurung buka, maka false
		if strings.Contains("}])", string([]rune(b)[0])) || strings.Contains("{[(", string([]rune(b)[len(b)-1])) {
			result = false
			break
		}
		//kalau kurung buka, masukkan ke stack
		if strings.Contains("{[(", string([]rune(b)[i])) {
			charStack = append(charStack, string([]rune(b)[i]))

			//kalau elemen terakhir stack bersesuaian dengan elemen yang sekarang, maka pop elemen terakhir
		} else if len(charStack) > 0 && charStack[len(charStack)-1] == prths[string([]rune(b)[i])] {
			fmt.Println("kesinii")
			charStack = charStack[:len(charStack)-1]
		} else {
			charStack = append(charStack, string([]rune(b)[i]))
		}
		result = false
		if len(charStack) < 1 {
			result = true
		}
		fmt.Println("charStack : ", charStack)
		//fmt.Println("charStack[len(charStack)-1]:", charStack[len(charStack)-1], ",prths[string([]rune(b)[i])]:", prths[string([]rune(b)[i])])
	}
	fmt.Println(charStack)

	return result
}
