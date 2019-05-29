package main

import "fmt"

type person struct {
	firstName string
	lastName string
	// contact contactInfo
	contactInfo
}

type contactInfo struct {
	email string
	zipCode int 
}

func main(){

	// var akram person
	// akram.firstName = "Akram"
	// akram.lastName = "Potrebin"

	// akram := person{firstName: "Akram", lastName: "potrebin"}
	// fmt.Println(akram)
	// fmt.Printf("%+v", akram)




	jim := person{
		firstName: "Jim",
		lastName: "Doe",
		contactInfo: contactInfo{
			email: "jim@gmail.com",
			zipCode: 94000,
		},
	}

	// jimPointer := &jim
	jim.updateName("Jimmy")
	jim.print() 

}



func (p person) print(){
	fmt.Printf("%+v",p)
}

func (pointerToPerson *person) updateName(newFirstName string){
	(*pointerToPerson).firstName = newFirstName
}
