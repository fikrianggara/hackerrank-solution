package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readInput() string {
	//read input
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	return text
}

type myList struct {
	data   []int
	max    int
	median float64
}

func convertStrToInt(s string) int {
	intConv, _ := strconv.Atoi(s)
	return intConv
}

func (a *myList) addElem(elem int) {
	//diasumsikan arraynya sudah diurutkan
	a.data = append(a.data, elem)
	//hanya disorting kalau elemen ynag dientri lebih kecil dari maksimal
	if elem < a.max {
		a.sortArray()
	} else {
		a.max = elem
	}
}

func (a *myList) removeElemByIndex(index int) {
	if index == 0 {
		a.data = a.data[1:]
	} else if index == len(a.data)-1 {
		a.data = a.data[:index]
	} else {
		temp1 := a.data[:index]
		temp2 := a.data[index+1:]
		a.data = append(temp1, temp2...)
	}
}

func (a *myList) removeElemByIndexes(indexes []int) {
	//ketika elemen diremove, maka elemen berikutnya akan bergeser ke kiri,
	//sehingga perlu dikurangi -1 untuk menyesuaikan perubahan index
	count := 0
	for _, item := range indexes {
		a.removeElemByIndex(item - count)
		count = count + 1
	}
}

func (a *myList) findMultipleElemIndex(elem int) []int {
	var result []int
	for index, item := range a.data {
		if item == elem {
			result = append(result, index)
		}
	}
	return result
}

func (a *myList) findElemIndex(elem int) (int, bool) {
	var result int
	found := false
	for index, item := range a.data {
		if item == elem {
			result = index
			found = true
			break
		}
	}
	return result, found
}

func (a *myList) CalculateMedian() {
	// fmt.Println(a)
	var n = len(a.data)
	if n == 0 {
		a.median = 0
	}
	if n == 1 {
		a.median = float64(a.data[0])
	} else if n == 2 {
		a.median = (float64(a.data[0]) + float64(a.data[1])) / 2.0

	} else {
		tengah := n / 2
		if n%2 == 0 {
			a.median = (float64(a.data[tengah])*1.0 + float64(a.data[tengah+1])) / 2.0
			// fmt.Println("genap, median =(", a.data[tengah], "+", a.data[tengah+1], ")/2=", a.median)
		} else {
			a.median = float64(a.data[tengah])
		}
	}

}

//bubble sort
func (a *myList) sortArray() {
	var (
		n      = len(a.data)
		sorted = false
	)
	if len(a.data) > 1 {

		for !sorted {
			swapped := false
			for i := 0; i < n-1; i++ {
				if a.data[i] > a.data[i+1] {
					a.data[i+1], a.data[i] = a.data[i], a.data[i+1]
					swapped = true
				}
			}
			if !swapped {
				sorted = true
			}
			n = n - 1
		}
	}
}

func initMedianUpdateSolutions() {

	fmt.Println("input text:")
	scanner := bufio.NewScanner(os.Stdin)
outerLoop:
	for {
		scanner.Scan()
		N := convertStrToInt(scanner.Text())
		var data = myList{}
		for i := 0; i < N; i++ {
			scanner.Scan()
			inputText := scanner.Text()
			inputTemp := strings.Split(inputText, " ")
			action := inputTemp[0]
			elem := convertStrToInt(inputTemp[1])

			if action == "a" {
				// fmt.Println("add ", elem)
				if i == 0 {
					data.max = elem
				}
				data.addElem(elem)
				data.CalculateMedian()
				// fmt.Println("data :", data.data, "median :", data.median)
				// fmt.Println("median : ", data.median)
				// median = data.getMedian()
				fmt.Println(data.median)
			} else if action == "r" {
				// fmt.Println("remove ", elem)
				idx, found := data.findElemIndex(elem)
				if !found {
					fmt.Println("Wrong!")
				} else {
					// if len(idxs) < 1 {
					// 	fmt.Println("Wrong!")
					// } else {
					//ini sudah terurut
					// fmt.Println("found ", len(idxs), "item. index:", idxs)
					data.removeElemByIndex(idx)
					// fmt.Println("data", data.data)
					// fmt.Println("data setelah dihapus", data)
					if len(data.data) < 1 {
						fmt.Println("Wrong!")
					} else {
						data.CalculateMedian()
						fmt.Println(data.median)
					}
					// fmt.Println("data :", data.data, "median :", data.median)
					// }

				}

			} else {
				fmt.Println("action not registered")
			}
		}
		break outerLoop
	}

}
