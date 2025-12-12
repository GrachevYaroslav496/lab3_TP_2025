package AbstractFactory

import "fmt"

type AsusMonitor struct {
	Size int
}

func (pc AsusMonitor) PrintDetails() {
	fmt.Printf("[ASUS] Size: %d\n", pc.Size)
}
