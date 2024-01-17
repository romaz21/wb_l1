package main

import "fmt"

type Human struct {
    Name string
    Age  int
}

type Action struct {
    Human
}

func (h *Human) SayHello() {
    fmt.Printf("Hello, my name is %s and I'm %d years old.\n", h.Name, h.Age)
}

func (a *Action) DoSomething() {
    fmt.Printf("%s is doing something...\n", a.Name)
}

func main() {
    actionInstance := Action{
        Human: Human{
            Name: "John",
            Age:  30,
        },
    }

    actionInstance.SayHello()
    actionInstance.DoSomething()
}

