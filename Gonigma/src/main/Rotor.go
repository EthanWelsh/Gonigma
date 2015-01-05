package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"unicode"
)

type Rotor struct {
	position int
	contacts []rune
}

// reads a rotor file from the given path and returns a rotor of that configuration
func (r Rotor) NewRotor(file string) Rotor {

	var err error

	message, err := ioutil.ReadFile(file) // Read the message from the message file

	if err != nil {
		fmt.Println("Error reading from file!!!")
		fmt.Println(err)
	}

	s := string(message[:])
	split := strings.Split(s, "\n")

	r.contacts = make([]rune, len(split))

	for i := range split {
		r.contacts[i] = rune([]rune(split[i])[0])
	}

	r.position = 0

	return r
}

// given a character, will translate that character and return the translated character
func (r *Rotor) Translate(c rune) (ret rune) {

	c = unicode.ToUpper(c)

	indexInRotor := (c - 'A' + rune(r.position)) % 26
	return r.contacts[indexInRotor]
}

// will set the rotor to a specific position
func (r *Rotor) SetToPosition(position int) {
	r.position = position
}

// rotates the rotor once and return if it is a kick (IE has it reached 26)
func (r *Rotor) RotateOnce() (isKick bool) {
	r.position = (r.position + 1) % 26
	return r.position == 0
}
