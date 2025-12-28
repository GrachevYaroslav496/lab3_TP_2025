package main

import "OtherCode/TP/Patterns/Builder"

func main() {
	asusCollector := Builder.GetCollector("Asus")
	applesCollector := Builder.GetCollector("Apple")

	factory := Builder.NewFactory(asusCollector)

	asusComputer := factory.CreateComputer()
	asusComputer.Print()

	factory.SetCollector(applesCollector)
	appleComputer := factory.CreateComputer()
	appleComputer.Print()
}
