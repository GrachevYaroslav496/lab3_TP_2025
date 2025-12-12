package AbstractFactory

type AsusFactory struct {
}

func (factory AsusFactory) GetComputer() Computer {
	return AsusComputer{
		Cpu:    8,
		Memory: 512,
	}
}

func (factory AsusFactory) GetMonitor() Monitor {
	return AsusMonitor{
		Size: 24,
	}
}
