package main

import (
	"fmt"
	"unicode"
)

type Machine struct {
	plugboard Plugboard

	r1 Rotor
	r2 Rotor
	r3 Rotor

	reflector Reflector
}

func (m Machine) newMachine(plug string, r1 string, r2 string, r3 string, reflect string) Machine {

	m.plugboard = m.plugboard.NewPlugboard(plug)

	m.r1 = m.r1.NewRotor(r1)
	m.r2 = m.r2.NewRotor(r2)
	m.r3 = m.r3.NewRotor(r3)

	m.reflector = m.reflector.NewReflector(reflect)
	return m
}

func (m *Machine) translate(c rune) (ret rune) {

	ret = m.plugboard.Translate(c)

	ret = m.r3.Translate(ret)
	ret = m.r2.Translate(ret)
	ret = m.r1.Translate(ret)

	ret = m.reflector.Translate(ret)

	ret = m.r1.ReverseTranslate(ret)
	ret = m.r2.ReverseTranslate(ret)
	ret = m.r3.ReverseTranslate(ret)

	ret = m.plugboard.Translate(ret)

	var isKick bool = m.r3.RotateOnce()

	if isKick {
		isKick = m.r2.RotateOnce()
		if isKick {
			m.r1.RotateOnce()
		}
	}

	return
}

func (m *Machine) reset() {
	m.r1.SetToPosition(0)
	m.r2.SetToPosition(0)
	m.r3.SetToPosition(0)
}

func (m *Machine) translateString(s string) (message string) {

	msg := []byte(s)

	for i := range msg {

		char := unicode.ToUpper(rune(msg[i]))

		if char >= 'A' && char <= 'Z' {
			message = message + string(m.translate(char))
		} else {
			message = message + string(char)
		}
	}

	return

}

func main() {

	var enigma Machine
	enigma = enigma.newMachine(
		"/Users/welshej/github/Enigma/ConfigFiles/p.plug",
		"/Users/welshej/github/Enigma/ConfigFiles/a.rotor",
		"/Users/welshej/github/Enigma/ConfigFiles/b.rotor",
		"/Users/welshej/github/Enigma/ConfigFiles/c.rotor",
		"/Users/welshej/github/Enigma/ConfigFiles/r.reflector")

	enigma.reset()

	fmt.Println(enigma.translateString("Hello World My Name Is Ethan"))

	enigma.reset()

	fmt.Println(enigma.translateString("EDHWZ AKMCT JQ LPOQ PK VWLMU"))

}
