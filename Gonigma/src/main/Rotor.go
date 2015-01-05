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

func (Rotor) NewRotor(file string) Rotor {

	var err error
	var r Rotor

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

func (r Rotor) translate(c rune) (ret rune) {

	c = unicode.ToUpper(c)
	indexInRotor := c - 'A'
	return r.contacts[indexInRotor]
}

/*
TODO IMPLEMENT
void setToPosition(int p)
void rotateOnce()
*/
