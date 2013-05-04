package main

import "fmt"

type Person struct {
    name string
    age int
    address string
}

func (person *Person) Name(name string){
    person.name = name
}

func (person *Person) Age(age int){
    person.age = age
}

func (person *Person) Address(add string) {
    person.address = add
}

func main() {
    p := Person{"dummy", 00, "god knows where"}
    fmt.Println(p)
    p.Name("Kiran")
    p.Age(23)
    p.Address("Tambaram East")
    fmt.Println(p)
}