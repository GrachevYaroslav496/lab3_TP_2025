package AbstractFactory

type AppleFactory struct {
}

func (factory AppleFactory) GetComputer() Computer {
	return AppleComputer{
		Cpu:    16,
		Memory: 1024,
	}
}

func (factory AppleFactory) GetMonitor() Monitor {
	return AppleMonitor{
		Size: 13,
	}
}
