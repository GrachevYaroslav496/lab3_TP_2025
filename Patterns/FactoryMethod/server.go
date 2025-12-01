package FactoryMethod

import "fmt"

type Server struct {
	Type   string
	Core   int
	Memory int
}

func NewServer() Computer {
	return Server{
		Type:   ServerType,
		Core:   64,
		Memory: 1024,
	}
}

func (s Server) GetType() string {
	return s.Type
}

func (s Server) PrintDetails() {
	fmt.Printf("Type: %s, Core: %d, Memory: %d\n", s.GetType(), s.Core, s.Memory)
}
