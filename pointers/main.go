package main

import "fmt"

type Person struct {
	firstName   string
	lastName    string
	contactInfo ContactInfo
}

type ContactInfo struct {
	email   string
	zipCode int
}

func (p *Person) updateName(newFirstName string) {
	(*p).firstName = newFirstName
}

func (p *Person) print() {
	fmt.Println(p)
}

func main() {
	alex := Person{"Alex", "Anderson",
		ContactInfo{
			email:   "email@alex.com",
			zipCode: 4826130,
		},
	}
	alex.print()

	alex.updateName("Alexey")
	alex.print()

	mySlice := []int{1, 2, 3, 4, 5}
	updateSlice(mySlice)
	fmt.Println(mySlice)
}

func updateSlice(s []int) {
	s[0] = 100
}
