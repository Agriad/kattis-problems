package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var inputAmount string

	fmt.Scanln(&inputAmount)

	var inputAmountString, _ = strconv.Atoi(inputAmount)

	var inputs []string
	var input string

	for i := 0; i < inputAmountString; i++ {
		fmt.Scanln(&input)
		inputs = append(inputs, input)
		fmt.Println("i")
	}

	// fmt.Println("Hello, World!")

	// var inputs [3]string = [3]string{"Sanna 1 16/03", "Simon 2 16/03", "Saga 3 14/10"}
	// var inputAmount int = 3

	// var inputs [10]string = [10]string{"Oden 78 03/12",
	// 	"Tor 132 14/05",
	// 	"Freja 10000 14/05",
	// 	"Loke 512 12/10",
	// 	"Hel 14 04/05",
	// 	"Fjorgynn 532 13/05",
	// 	"Hildegun 500 13/05",
	// 	"Vindsval 17 03/12",
	// 	"Snotra 20 04/05",
	// 	"Kvaser 420 03/12"}
	// var inputAmount int = 10

	var parsedInput []string

	for i := 0; i < inputAmountString; i++ {
		var tempMain []string
		var toAdd int = 1
		tempMain = strings.Split(inputs[i], " ")

		for j := 0; j < inputAmountString; j++ {
			var temp []string
			temp = strings.Split(inputs[j], " ")

			if tempMain[2] == temp[2] {
				var tempMainInt, _ = strconv.Atoi(tempMain[1])
				var tempInt, _ = strconv.Atoi(temp[1])

				if tempMainInt < tempInt {
					toAdd = 0
				}
			}
		}

		if toAdd == 1 {
			parsedInput = append(parsedInput, string(tempMain[0]))
		}
	}

	fmt.Println(len(parsedInput))

	for i := 0; i < len(parsedInput); i++ {
		fmt.Println(parsedInput[i])
	}
}

// https://open.kattis.com/problems/fodelsedagsmemorisering
// https://go.dev/doc/tutorial/getting-started
// https://stackoverflow.com/questions/20895552/how-can-i-read-from-standard-input-in-the-console
