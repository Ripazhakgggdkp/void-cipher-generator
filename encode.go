package main

import (
	"fmt"
	"strconv"
)

func letterToGrid(s rune) [8]bool {
	if string(s) == "/" {
		return [8]bool{}
	}

	value := int(s) - 64
	binary := []rune(fmt.Sprintf("%08b", rune(value)))
	return [8]bool{true, false, false, conv(binary[3]), conv(binary[4]), conv(binary[5]), conv(binary[7]), conv(binary[6])}
}

func sentenceToCipher(s string) (sentence [][8]bool) {
	for _, s := range s {
		sentence = append(sentence, letterToGrid(s))
	}
	return
}

func conv(s rune) bool {
	result, _ := strconv.Atoi(string(s))
	return result == 1
}
