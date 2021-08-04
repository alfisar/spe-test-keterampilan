package main

import (
	"fmt"
	"math"
	"strconv"
)

func narcisstic(data string, inputData int) (result bool, err error) {
	lenInput := len(data)

	count := 0

	for k := 0; k < lenInput; k++ {
		tmp, err := strconv.Atoi(string(data[k]))
		if err != nil {
			return result, err
		}

		count = count + int(math.Pow(float64(tmp), float64(lenInput)))
	}

	result = count == inputData
	return
}

func main() {
	var dataInput int

	fmt.Print("Please Input The Number : ")
	fmt.Scanln(&dataInput)

	convert := strconv.Itoa(dataInput)

	result, err := narcisstic(convert, dataInput)
	if err != nil {
		fmt.Println("Input Must Be Interger")
	} else {
		fmt.Println(result)
	}

}
