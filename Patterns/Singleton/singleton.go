package main

import "fmt"

var singleton *Singleton

type Singleton struct {
	Type string
}

func init() {
	fmt.Println("init: базовая иницализация")
	singleton = &Singleton{Type: "Одиночка"}
}

func main() {
	for i := 0; i < 3; i++ {
		singleton = newSingleton(singleton, "create newType")
	}
}

func newSingleton(item *Singleton, typeName string) *Singleton {
	if item == nil {
		return &Singleton{typeName}
	}
	fmt.Printf("Type %s - уже создан\n", item.Type)
	return item
}
