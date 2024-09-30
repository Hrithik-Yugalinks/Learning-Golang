package main

import (
	"fmt"
	"os"
	"strconv"
)

// func main() {
// 	arguments := os.Args

// 	if len(arguments) == 1 {
// 		fmt.Println("Provide an argument")
// 		return
// 	}

// 	var total, nInts, nFloats int
// 	invalid := make([]string, 0)

// 	for _, k := range arguments[1:] {
// 		_, err := strconv.Atoi(k)

// 		if err == nil {
// 			total++
// 			nInts++
// 			continue
// 		}

// 		_, err1 := strconv.ParseFloat(k, 64)

// 		if err1 == nil {
// 			total++
// 			nFloats++
// 			continue
// 		}

// 		invalid = append(invalid, k)

// 	}

// 	fmt.Println("#read:", total, "#ints:", nInts, "#floats:", nFloats)

// 	if len(invalid) > total {
// 		fmt.Println("Too much invalid input:", len(invalid))

// 		for _, s := range invalid {
// 			fmt.Println(s)
// 		}
// 	}
// }

func main() {
	arguments := os.Args

	if len(arguments) == 1 {
		fmt.Println("Provide an argument")
		return
	}

	var total, nInts, nFloats int

	invalid := make([]string, 0)

	for _, value := range arguments[1:] {

		_, err := strconv.Atoi(value)

		if err == nil {
			total++
			nInts++
			continue
		}

		_, err1 := strconv.ParseFloat(value, 64)

		if err1 == nil {
			total++
			nFloats++
			continue
		}

		invalid = append(invalid, value)
	}

	fmt.Println("#read:", total, "#ints:", nInts, "#floats:", nFloats)

	if len(invalid) > total {
		fmt.Println("Too much invalid input:", len(invalid))

		for _, value := range invalid {
			fmt.Println(value)
		}
	}

}
