package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"unicode"
)

type Reflector struct {
	plugs []rune
}

// reads a reflector file from the given path and returns a rotor of that configuration
func (r Reflector) NewReflector(file string) Reflector {

	var err error

	message, err := ioutil.ReadFile(file) // Read the message from the message file

	if err != nil {
		fmt.Println("Error reading from file!!!")
		fmt.Println(err)
	}

	s := string(message[:])
	split := strings.Split(s, "\n")

	r.plugs = make([]rune, len(split))

	for i := range split {
		r.plugs[i] = rune([]rune(split[i])[0])
	}

	return r
}

// given a character, will translate that character and return the translated character
func (r *Reflector) Translate(c rune) (ret rune) {

	c = unicode.ToUpper(c)

	indexInRotor := c - 'A'
	return r.plugs[indexInRotor]
}
