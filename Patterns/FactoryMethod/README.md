Для проверки работы FactoryMethod написало в `main.go` вызов пакета, в `types` указал типы для теста кода.
```main.go
package main

import fm "OtherCode/TP/Patterns/FactoryMethod"

var types = []string{fm.ServerType, fm.MobileType, fm.PersonComputerType, "пупупу"}

func main() {
for _, typeName := range types {
computer := fm.New(typeName)
if computer == nil {
continue
}
computer.PrintDetails()
}
}
```