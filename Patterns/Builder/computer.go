package Builder

import "fmt"

type Computer struct {
	Core   int
	Brand  string
	Memory int
}

func (pc *Computer) Print() {
	fmt.Printf("[%s]: Core: %d; Memory: %d\n", pc.Brand, pc.Core, pc.Memory)
}
