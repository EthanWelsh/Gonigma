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

	message, err := ioutil.ReadFile(file)

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

// given a letter, will return the letter which translates into this letter
func (r *Rotor) ReverseTranslate(c rune) (ret rune) {

	var i rune

	for i = 0; i < 26; i++ {

		if r.contacts[i] == c {
			ret = i + 'A' - rune(r.position)

			if ret < 'A' {
				ret = 'Z' - ('A' - ret) + 1
			}
		}
	}

	return
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
