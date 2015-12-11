package main

import "log"

func main() {
	p := new(Person)
	p.Name = "Hector"
	p.Talk() 
	
	a := new(Android)
	a.Name = "ARG3000"
	a.Talk() //we can call any Person methods directly on Android. The is-a relationship works this way intuitively: People can talk, an android is a person, therefore an android can talk.
}
// People can talk...
func (randomName *Person) Talk() {
	log.Println("Hi, my name is", randomName.Name)
}
type Person struct {
	Name string
}
type Android struct {
	Person //embedded type, also known as anonymous field. We use the type Person and don't give it a name.
	Model string
}