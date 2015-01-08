package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"unicode"
)

type Reflector struct {
	slots []rune
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

	if len(split)*2 != 26 {
		fmt.Println("Reflector file formatted incorrectly. Exiting program!")
		os.Exit(0)
	}

	r.slots = make([]rune, 26)

	for i := range split {

		pair := strings.Split(split[i], " ")

		first := pair[0]
		second := pair[1]

		firstRune := ([]rune(first))[0]
		secondRune := ([]rune(second))[0]

		r.slots[firstRune-'A'] = secondRune
		r.slots[secondRune-'A'] = firstRune
	}

	return r
}

// given a character, will translate that character and return the translated character
func (r *Reflector) Translate(c rune) (ret rune) {

	c = unicode.ToUpper(c)
	index := c - 'A'
	return r.slots[index]
}
