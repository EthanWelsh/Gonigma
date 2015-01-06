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

	p.plugs = make([]rune, len(split))

	for i := range split {
		p.plugs[i] = rune([]rune(split[i])[0])
	}

	return p
}

// given a character, will translate that character and return the translated character
func (p *Plugboard) Translate(c rune) (ret rune) {

	c = unicode.ToUpper(c)

	indexInRotor := c - 'A'
	return p.plugs[indexInRotor]
}
