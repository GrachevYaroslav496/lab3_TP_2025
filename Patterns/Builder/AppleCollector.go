package Builder

type AppleCollector struct {
	Core   int
	Brand  string
	Memory int
}

func (c *AppleCollector) SetCore() {
	c.Core = 8
}

func (c *AppleCollector) SetBrand() {
	c.Brand = "Apple"
}

func (c *AppleCollector) SetMemory() {
	c.Memory = 512
}

func (c *AppleCollector) GetComputer() Computer {
	return Computer{
		Core:   c.Core,
		Brand:  c.Brand,
		Memory: c.Memory,
	}
}
