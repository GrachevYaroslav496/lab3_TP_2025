package AbstractFactory

import "fmt"

type AppleComputer struct {
	Cpu    int
	Memory int
}

func (pc AppleComputer) PrintDetails() {
	fmt.Printf("[APPLE] Cpu: %d, Memory: %d\n", pc.Cpu, pc.Memory)
}
