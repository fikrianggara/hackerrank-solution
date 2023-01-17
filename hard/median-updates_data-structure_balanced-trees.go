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
	median float32
}

func convertStrToInt(s string) int {
	intConv, _ := strconv.Atoi(s)
	return intConv
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

func (a *myList) CalculateMedian() {
	// fmt.Println(a)
	var n = len(a.data)
	if n <= 1 {
		a.median = float32(a.data[0])
	} else if n == 2 {
		a.median = float32((a.data[0] + a.data[1]) / 2.0)

	} else {
		tengah := n / 2
		if n%2 == 0 {
			fmt.Println("genap, median =(", a.data[tengah], "+", a.data[tengah+1], ")/2=")
			a.median = float32((a.data[tengah] + a.data[tengah+1]) / 2.0)
			fmt.Println("genap, median =(", a.data[tengah], "+", a.data[tengah+1], ")/2=", a.median)
		} else {
			a.median = float32(a.data[tengah])
		}
	}

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

func (a *myList) removeElem(elem int) {

}

func (a *myList) findElem(elem int) {

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
				fmt.Println("add ", elem)
				if i == 0 {
					data.max = elem
				}
				data.addElem(elem)
				data.CalculateMedian()
				fmt.Println("data :", data.data, "median :", data.median)
				// median = data.getMedian()
			} else if action == "r" {
				fmt.Println("remove ", elem)
				data.removeElem(elem)
				if len(data.data) < 1 {
					fmt.Println("Wrong!")
				}
			} else {
				fmt.Println("action not registered")
			}
		}
		break outerLoop
	}

}
