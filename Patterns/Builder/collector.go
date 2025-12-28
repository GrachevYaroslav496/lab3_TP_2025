package Builder

const (
	AppleCollectorType = "Apple"
	AsusCollectorType  = "Asus"
)

type Collector interface {
	SetCore()
	SetBrand()
	SetMemory()
	GetComputer() Computer
}

func GetCollector(collectorType string) Collector {
	switch collectorType {
	default:
		return nil
	case AppleCollectorType:
		return &AppleCollector{}
	case AsusCollectorType:
		return &AsusCollector{}

	}
}
