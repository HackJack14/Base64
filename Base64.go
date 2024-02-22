package main

import (
	"fmt"
	"os"
)

var charSet = [64]string{
	"A", "B", "C", "D", "E", "F", "G", "H", "I", "J",
	"K", "L", "M", "N", "O", "P", "Q", "R", "S", "T",
	"U", "V", "W", "X", "Y", "Z", "a", "b", "c", "d",
	"e", "f", "g", "h", "i", "j", "k", "l", "m", "n",
	"o", "p", "q", "r", "s", "t", "u", "v", "w", "x",
	"y", "z", "0", "1", "2", "3", "4", "5", "6", "7",
	"8", "9", "+", "/"}

var isDecoding bool = false
var inRunes []rune
var output string = ""

func parseInput() bool {
	if len(os.Args) > 2 {
		if os.Args[1] == "-d" {
			isDecoding = true
		} else if os.Args[1] == "-e" {
			isDecoding = false
		} else {
			fmt.Println("Option not supported")
			return false
		}
		input := os.Args[2]
		for i := 0; i < len(input); i++ {
			inRunes = append(inRunes, rune(input[i]))
		}
		return true
	} else {
		fmt.Println("Not enough arguments")
		return false
	}
}

func encode() {
	numRunes := len(inRunes)
	for i := 0; i < numRunes/3; i++ {
		currRune := i * 3
		output = output + charSet[inRunes[currRune]>>2]
		output = output + charSet[((inRunes[currRune]&0b11)<<4)+(inRunes[currRune+1])>>4]
		output = output + charSet[((inRunes[currRune+1]&0b1111)<<2)+(inRunes[currRune+2])>>6]
		output = output + charSet[inRunes[currRune+2]&0b111111]
	}

	remainder := numRunes % 3

	switch remainder {
	case 1:
		output = output + charSet[inRunes[numRunes-1]>>2]
		output = output + charSet[((inRunes[numRunes-1]&0b11)<<4)]
		output = output + "=="
	case 2:
		output = output + charSet[inRunes[numRunes-2]>>2]
		output = output + charSet[((inRunes[numRunes-2]&0b11)<<4)+(inRunes[numRunes-1])>>4]
		output = output + charSet[((inRunes[numRunes-1]&0b1111)<<2)]
		output = output + "="
	}

	fmt.Println(output)
}

func decode() {

}

func main() {
	if parseInput() == true {
		encode()
	}
}
