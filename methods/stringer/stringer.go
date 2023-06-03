package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v 歳)", p.Name, p.Age)
}

func main() {
	a := Person{"太郎", 20}
	b := Person{"花子", 21}
	fmt.Println(a, b)
}
