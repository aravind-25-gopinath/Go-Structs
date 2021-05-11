package main

import "fmt"

//struct
type contactInfo struct {
	email   string
	zipcode int
}

//struct
type Person struct {
	firstName string      //firstname attribute of type string
	lastName  string      //lastname attribute of type string
	contact   contactInfo //embedded struct (using a previulsy created struct (contactinfo) as an attribute in the Person struct)
}

func main() {
	//alex := Person{"Alex", "Anderson"} //one way of initializing a struct

	//alex := Person{firstName: "Alex", lastName: "Anderson"} //another way of initialzing struct (easier to update struct fields)
	//fmt.Println(alex)

	//another way of initialzing struct
	/*var alex Person
	alex.firstName = "Alex"
	alex.lastName = "Anderson"
	fmt.Println(alex)*/

	//initializing an embedded struct
	jim := Person{firstName: "Jim", lastName: "Party", contact: contactInfo{email: "jim@gmail.com", zipcode: 94000}}
	//fmt.Println(jim)
	jim.print()
}

//reciever function to print
func (p Person) print() {
	fmt.Println(p)
}
