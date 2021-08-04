package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	var dataInput int

	fmt.Print("Please Input Len Array : ")
	fmt.Scanln(&dataInput)

	if dataInput < 3 {
		log.Fatalln("Your Input Must Be Length Minimum 3")
	}

	tmp := make([]int, 0)
	odd, even := 0, 0
	oddNumber, evenNumber := 0, 0
	for k := 0; k < dataInput; k++ {
		var dataInputElmnt int
		convert := strconv.Itoa(k)
		fmt.Print("Please Input data " + convert + " : ")
		fmt.Scanln(&dataInputElmnt)
		tmp = append(tmp, dataInputElmnt)
		if dataInputElmnt%2 == 0 {
			even = even + 1
			evenNumber = dataInputElmnt
		} else {
			odd = odd + 1
			oddNumber = dataInputElmnt
		}
	}

	if dataInput-1 == even && odd == 1 {
		convert := strconv.Itoa(oddNumber)
		fmt.Println("Should return : " + convert + "(the only odd number)")
	} else if dataInput-1 == odd && even == 1 {
		convert := strconv.Itoa(evenNumber)
		fmt.Println("Should return : " + convert + "(the only even number)")
	} else if dataInput == odd {
		fmt.Println("false (all odd integer, no outlier)")
	} else if dataInput == even {
		fmt.Println("false (all even integer, no outlier)")
	} else {
		fmt.Println("false (all integer, no outlier)")
	}
}
