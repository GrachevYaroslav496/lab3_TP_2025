package Builder

type AsusCollector struct {
	Core   int
	Brand  string
	Memory int
}

func (c *AsusCollector) SetCore() {
	c.Core = 4
}

func (c *AsusCollector) SetBrand() {
	c.Brand = "Asus"
}

func (c *AsusCollector) SetMemory() {
	c.Memory = 1024
}

func (c *AsusCollector) GetComputer() Computer {
	return Computer{
		Core:   c.Core,
		Brand:  c.Brand,
		Memory: c.Memory,
	}
}
