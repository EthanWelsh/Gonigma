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
	plugPairs := strings.Split(s, "\n")

	p.plugs = make([]rune, 26)

	for i := range plugPairs {
		pair := strings.Split(plugPairs[i], " ")

		first := pair[0]
		second := pair[1]

		firstRune := ([]rune(first))[0]
		secondRune := ([]rune(second))[0]

		p.plugs[firstRune-'A'] = secondRune
		p.plugs[secondRune-'A'] = firstRune
	}

	for i := range p.plugs {
		if p.plugs[i] == 0 { // If a letter isn't mapped to another on the plugboard, map it to itself
			p.plugs[i] = rune(i) + 'A'
		}
	}
	return p
}

// given a character, will translate that character and return the translated character
func (p *Plugboard) Translate(c rune) (ret rune) {

	c = unicode.ToUpper(c)

	index := c - 'A'
	return p.plugs[index]
}
