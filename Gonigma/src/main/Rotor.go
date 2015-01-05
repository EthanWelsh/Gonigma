package main

import (
	"fmt"
	"io/ioutil"
)

type Rotor struct {
	position int
	contacts []byte
}

func (Rotor) newRotor(file string) Rotor {

	var err error

	message, err := ioutil.ReadFile(file) // Read the message from the message file

	if err != nil {
		print("Error reading from file!!!")
	}

	fmt.Println(message)

	var r Rotor

	return r
}

/*Rotor(const char *rotor_file);
char translate(char c);
void setToPosition(int p);
void rotateOnce();

private:
int position;
char contacts[];*/
