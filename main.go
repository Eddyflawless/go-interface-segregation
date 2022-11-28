package main

import (
	"fmt"
)

type Warrior interface {
	Punch()
	Kick()
}

type SuperSaiyan interface {
	Transform()
}

type Warriors []Warrior //  Derived type

// MrSatan implementation

type MrSatan struct {
	//properties
	hairColor string
}

func NewMrSatan() *MrSatan {
	return &MrSatan{hairColor: "black"}
}

func (m MrSatan) Punch() {
	fmt.Println("Mr Satan Punchs")
}

func (m MrSatan) Kick() {
	fmt.Println("Mr Satan Kicks")
}

func (m MrSatan) Transform() {
	// do nohting here
	fmt.Errorf("%v /n", "No implementation..")
}

type Goku struct {
	hairColor string
}

func NewGoku() *Goku {

	return &Goku{"blue"}
}

func (g Goku) Punch() {
	fmt.Println("Goku Punchs")
}

func (g Goku) Kick() {

	fmt.Println("GOku kicks")
}

func (g Goku) Transform() {

	defer func() {
		// if err := recover(); err != nil {
		// 	fmt.Println("[OTHER] recovered from panic: 2")
		// }
	}()

	defer func() {

		// if r := recover(); r != nil {
		// 	fmt.Printf("[OTHER] Recovered from a panic 3 \n")
		// }
		fmt.Println("GOku transforms into a super saiyan -->")

	}()

	fmt.Println("Goku screams and charges up")
	{
		panic("panic because why not")
	}

	fmt.Println("all done")
}

func executeWihoutISP(warriors Warriors) {

	for _, warrior := range warriors {
		warrior.Punch()
		warrior.Kick()

		// check if each warrior is a super saiyan before executing Transform()
		if superSaiyan, ok := warrior.(SuperSaiyan); ok {
			superSaiyan.Transform()
		}
	}
}

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("[MAIN] recovered from panic: 1") // handles any uncaught errors
		}
	}()
	fmt.Println("Interface segregation...")
	warriors := Warriors{}
	warriors = append(warriors, NewGoku())
	warriors = append(warriors, NewMrSatan())

	//
	executeWihoutISP(warriors)

}
