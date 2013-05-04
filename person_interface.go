package main

import "fmt"

type Person interface{
    greet() string
}

type Me string

func (me Me) greet() string {
    return string(me) + ": Hi, good morning."
}

func main() {
    var person Person
    person = Me("Kiran")
    fmt.Println(person.greet())
}