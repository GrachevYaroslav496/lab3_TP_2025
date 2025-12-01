package FactoryMethod

import "fmt"

type Mobile struct {
	Type   string
	Core   int
	Memory int
}

func NewMobile() Computer {
	return Mobile{
		Type:   MobileType,
		Core:   64,
		Memory: 1024,
	}
}

func (m Mobile) GetType() string {
	return m.Type
}

func (m Mobile) PrintDetails() {
	fmt.Printf("Type: %s, Core: %d, Memory: %d\n", m.GetType(), m.Core, m.Memory)
}
