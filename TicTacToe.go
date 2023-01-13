package main

import (
	"fmt"
	"strconv"
)

var fields [9]int
var input int

func printFields(fields [9]int) {
	var chars [9]string
	for i := 0; i < 9; i++ {
		if fields[i] == 0 {
			chars[i] = strconv.Itoa(i + 1)
		} else if fields[i] == 1 {
			chars[i] = "X"
		} else {
			chars[i] = "O"
		}
	}
	fmt.Println(" " + chars[0] + " | " + chars[1] + " | " + chars[2] + " ")
	fmt.Println("---+---+---")
	fmt.Println(" " + chars[3] + " | " + chars[4] + " | " + chars[5] + " ")
	fmt.Println("---+---+---")
	fmt.Println(" " + chars[6] + " | " + chars[7] + " | " + chars[8] + " ")
	fmt.Println()
}

func win(fields [9]int) bool {
	if fields[0] == fields[1] && fields[1] == fields[2] && fields[0] != 0 {
		return true
	} else if fields[3] == fields[4] && fields[4] == fields[5] && fields[3] != 0 {
		return true
	} else if fields[6] == fields[7] && fields[7] == fields[8] && fields[6] != 0 {
		return true
	} else if fields[0] == fields[3] && fields[3] == fields[6] && fields[0] != 0 {
		return true
	} else if fields[1] == fields[4] && fields[4] == fields[7] && fields[1] != 0 {
		return true
	} else if fields[2] == fields[5] && fields[5] == fields[8] && fields[2] != 0 {
		return true
	} else if fields[0] == fields[4] && fields[4] == fields[8] && fields[0] != 0 {
		return true
	} else if fields[2] == fields[4] && fields[4] == fields[6] && fields[2] != 0 {
		return true
	}
	return false
}

func main() {
	for true {
		printFields(fields)
		fmt.Print("Player 1: ")
		fmt.Scan(&input)
		fmt.Println()
		if input > 0 && input <= 9 && fields[input-1] == 0 {
			fields[input-1] = 1
		}
		if win(fields) {
			printFields(fields)
			fmt.Println("Player 1 wins!")
			break
		}
		for i := 0; i < 9; i++ {
			if fields[i] == 0 {
				break
			} else if i == 8 {
				printFields(fields)
				fmt.Println("Draw!")
				return
			}
		}
		printFields(fields)
		fmt.Print("Player 2: ")
		fmt.Scan(&input)
		fmt.Println()
		if input > 0 && input <= 9 && fields[input-1] == 0 {
			fields[input-1] = 2
		}
		if win(fields) {
			printFields(fields)
			fmt.Println("Player 2 wins!")
			break
		}
	}
}
