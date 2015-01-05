package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type Rotor struct {
	position int
	contacts []byte
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

	r.contacts = make([]byte, len(split))

	for i := range split {
		r.contacts[i] = byte([]rune(split[i])[0])
	}

	r.position = 0

	return r
}

/*
TODO IMPLEMENT
char translate(char c)
void setToPosition(int p)
void rotateOnce()
*/
