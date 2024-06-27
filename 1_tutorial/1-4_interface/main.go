package main

import (
    "fmt"
)

func main() {
    // インターフェース
    fmt.Println("=== interface test ===")

    var h1 Animal = Human{"Tom"}
    h2 := Human{"Bob"}
    TalkTo(h1) // h1はAnimal型で宣言しているためOK
    TalkTo(h2) // h2はHuman型で宣言しており、Animalインターフェースに沿った型なのでOK

    var d Animal = Dog{"Hachi"}
    TalkTo(d)

    c := Cat{"Jiji"}
    TalkTo(c)

    //var p1 Animal = Plant{"Apple"} // PlantはAnimalインターフェースに沿う型ではないため不可

    //p2 := Plant{"Orange"}
    //TalkTo(p2) // p2はAnimalインターフェースに沿っていないため不可
}

// Animalインターフェース
// Speakメソッドを持つ必要がある
type Animal interface {
    Speak()
}

func TalkTo(a Animal) {
    fmt.Println("Hello, what's your name?")
    a.Speak()
    fmt.Println("")
}

// Human型はSpeakメソッドを持つ(=Animalインターフェースに沿っている)
type Human struct {
    name string
}
func (h Human) Speak() {
    fmt.Println("Hello, my name is", h.name, ".")
}

// Dog型はSpeakメソッドを持つ(=Animalインターフェースに沿っている)
type Dog struct {
    name string
}
func (d Dog) Speak() {
    fmt.Println("Bow wow")
}

// Cat型はSpeakメソッドを持つ(=Animalインターフェースに沿っている)
type Cat struct {
    name string
}
func (c Cat) Speak() {
    fmt.Println("Meow")
}

// Plant型はSpeakメソッドを持たない(=Animalインターフェースに沿っていない)
type Plant struct {
    name string
}
