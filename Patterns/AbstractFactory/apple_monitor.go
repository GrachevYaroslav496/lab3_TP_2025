package AbstractFactory

import "fmt"

type AppleMonitor struct {
	Size int
}

func (pc AppleMonitor) PrintDetails() {
	fmt.Printf("[APPLE] Size: %d\n", pc.Size)
}
