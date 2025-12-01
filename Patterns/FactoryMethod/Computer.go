package FactoryMethod

import "fmt"

const (
	ServerType         = "server"
	PersonComputerType = "personal"
	MobileType         = "mobile"
)

type Computer interface {
	GetType() string
	PrintDetails()
}

func New(typeName string) Computer {
	switch typeName {
	case ServerType:
		return NewServer()
	case PersonComputerType:
		return NewPersonal()
	case MobileType:
		return NewMobile()
	default:
		fmt.Printf("Неизвестный тип: %s\n", typeName)
		return nil
	}
}
