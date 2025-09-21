package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

type Employee struct {
	Person

	EmployeeID string
}

func (this *Employee) PrintInfo() {
	fmt.Println("Name=", this.Name, ",Age=", this.Age, ",EmployeeID", this.EmployeeID)
}

func main() {
	em := &Employee{

		Person: Person{
			Name: "test",
			Age:  123,
		},
		EmployeeID: "3213",
	}
	em.PrintInfo()
}
