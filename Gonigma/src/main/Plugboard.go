package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"unicode"
)

type Plugboard struct {
	plugs []rune
}

// reads a reflector file from the given path and returns a rotor of that configuration
func (p Plugboard) NewPlugboard(file string) Plugboard {

	var err error

	message, err := ioutil.ReadFile(file) // Read the message from the message file

	if err != nil {
		fmt.Println("Error reading from file!!!")
		fmt.Println(err)
	}

	s := string(message[:])
	split := strings.Split(s, "\n")

	p.plugs = make([]rune, len(split)*2)

	for i := range split {

		pair := strings.Split(split[i], " ")

		first := pair[0]
		second := pair[1]

		firstRune := ([]rune(first))[0]
		secondRune := ([]rune(second))[0]

		p.plugs[firstRune-'A'] = secondRune
		p.plugs[secondRune-'A'] = firstRune
	}

	return p
}

// given a character, will translate that character and return the translated character
func (p *Plugboard) Translate(c rune) (ret rune) {

	c = unicode.ToUpper(c)

	indexInRotor := c - 'A'
	return p.plugs[indexInRotor]
}
