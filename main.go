package main

import (
	"fmt"
	"strings"
)

type Employee struct {
	name     string
	age      int
	division string
}

type Bank struct {
	bank     string
	employee Employee
}

func main() {
	var firstPerson string = "Arif"
	var secondPerson *string = &firstPerson

	fmt.Println("firstPerson (value) : ", firstPerson)
	fmt.Println("firstPerson (address) : ", &firstPerson)

	fmt.Println("secondPerson (value) : ", *secondPerson)
	fmt.Println("secondPerson (address) : ", secondPerson)

	*secondPerson = "Muspita"

	fmt.Println("firstPerson (value) : ", firstPerson)
	fmt.Println("firstPerson (address) : ", &firstPerson)

	fmt.Println("secondPerson (value) : ", *secondPerson)
	fmt.Println("secondPerson (address) : ", secondPerson)

	var employee1 = Employee{name: "Arif Muspita", age: 23, division: "Gas"}

	var employee2 *Employee = &employee1

	fmt.Println("Employee 1 name : ", employee1.name)
	fmt.Println("Employee 2 name : ", employee2.name)

	fmt.Println(strings.Repeat("#", 20))

	employee2.name = "Muspita Arif "

	fmt.Println("Employee 1 name : ", employee1.name)
	fmt.Println("Employee 2 name : ", employee2.name)

	var bank = Bank{}

	bank.bank = "BCA"
	bank.employee.age = 34
	bank.employee.name = "AM"
	bank.employee.division = "Pindah"

	fmt.Printf("%+v\n", bank)

	var employee4 = []Employee{
		{name: "Arif", age: 22, division: "A"},
		{name: "Muspita", age: 33, division: "C"},
		{name: "AM", age: 11, division: "B"},
	}

	for _, v := range employee4 {
		fmt.Printf("%+v\n", v)
	}

	var employee5 = Employee{name: "Gue", age: 22, division: "Z"}

	fmt.Println(employee5.Introduce("Welcome"))

	var employee6 = Employee{name: "Loe", age: 12, division: "ZZZ"}

	employee6.ChangeName1()
	fmt.Println("Change name with method", employee6.name)

	employee6.ChangeName2()
	fmt.Println("Change name with method pointer", employee6.name)

}

func (p Employee) Introduce(msg string) string {
	return fmt.Sprintf("%s My name is %s and i'm %d years old. Hello from method", msg, p.name, p.age)
}

func (p Employee) ChangeName1() {
	p.name = "Gas"
}

func (p *Employee) ChangeName2() {
	p.name = "Brumm"
}
