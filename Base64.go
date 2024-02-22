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
var outRunes []rune

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
	first := charSet[inRunes[0]>>2]
	second := charSet[((inRunes[0]&0b11)<<4)+(inRunes[1])>>4]
	third := charSet[((inRunes[1]&0b1111)<<2)+(inRunes[2])>>6]
	fourth := charSet[inRunes[2]&0b111111]
	fmt.Println(first + second + third + fourth)
}

func decode() {

}

func main() {
	if parseInput() == true {
		encode()
	}
}
