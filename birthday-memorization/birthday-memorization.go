package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var inputAmount string

	fmt.Scanln(&inputAmount)

	var inputAmountString, _ = strconv.Atoi(inputAmount)

	var inputs []string
	var input string

	reader := bufio.NewReader(os.Stdin)

	for i := 0; i < inputAmountString; i++ {
		input, _ = reader.ReadString('\n')
		inputs = append(inputs, input)
	}

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

	sort.Strings(parsedInput)

	fmt.Println(len(parsedInput))

	for i := 0; i < len(parsedInput); i++ {
		fmt.Println(parsedInput[i])
	}
}

// test input

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

// https://open.kattis.com/problems/fodelsedagsmemorisering
// https://go.dev/doc/tutorial/getting-started
// https://stackoverflow.com/questions/20895552/how-can-i-read-from-standard-input-in-the-console
