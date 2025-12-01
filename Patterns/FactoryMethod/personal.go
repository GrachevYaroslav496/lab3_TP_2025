package FactoryMethod

import "fmt"

type Personal struct {
	Type   string
	Core   int
	Memory int
}

func NewPersonal() Computer {
	return Personal{
		Type:   PersonComputerType,
		Core:   64,
		Memory: 1024,
	}
}

func (pc Personal) GetType() string {
	return pc.Type
}

func (pc Personal) PrintDetails() {
	fmt.Printf("Type: %s, Core: %d, Memory: %d\n", pc.GetType(), pc.Core, pc.Memory)
}
