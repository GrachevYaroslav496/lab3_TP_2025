package Builder

func main() {
	asusCollector := GetCollector("Asus")
	applesCollector := GetCollector("Apple")

	factory := NewFactory(asusCollector)

	asusComputer := factory.CreateComputer()
	asusComputer.Print()

	factory.SetCollector(applesCollector)
	appleComputer := factory.CreateComputer()
	appleComputer.Print()
}
