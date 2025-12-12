package AbstractFactory

import (
	"fmt"
)

type Factory interface {
	GetMonitor() Monitor
	GetComputer() Computer
}

func GetFactory(brend string) Factory {
	switch brend {
	default:
		fmt.Println("Неизвестный бренд")
		return nil
	case Apple:
		return &AppleFactory{}
	case Asus:
		return &AsusFactory{}
	}
}
