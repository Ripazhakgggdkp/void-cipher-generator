package main

import (
	"fmt"
	"strconv"
)

type State int

const (
	Blank State = iota
	Off
	On
)

func letterToGrid(s rune) [3][3]State {
	if string(s) == "/" {
		return [3][3]State{}
	}

	value := int(s) - 64
	binary := []rune(fmt.Sprintf("%08b", rune(value)))
	return [3][3]State{
		{On, Off, Off},
		{conv(binary[3]), Blank, conv(binary[4])},
		{conv(binary[5]), conv(binary[7]), conv(binary[6])},
	}
}

func sentenceToCipher(s string) (sentence [][3][3]State) {
	for _, s := range s {
		sentence = append(sentence, letterToGrid(s))
	}
	return
}

func conv(s rune) State {
	result, _ := strconv.Atoi(string(s))

	if result == 1 {
		return On
	} else {
		return Off
	}
}
