package main

import "fmt"

type user struct {
	name string
	age  uint8
}

func (user user) save() {
	fmt.Printf("Executing save() on %s \n", user.name)
}

func (user user) canDrink() bool {
	if user.age > 20 {
		return true
	} else {
		return false
	}
}

func (user *user) addAge() {
	user.age++
} // to alter an instanced object, use it's pointer in order to persist changes outside of the scope of the method

func main() {
	fmt.Println("Methods")
	// Methods, unlike functions, are associated with specific objects, like classes, structs, etc.
	user1 := user{"Mad", 30}
	user1.save()

	user2 := user{"Dog", 20}
	user2.save()

	fmt.Println(user1.canDrink(), user2.canDrink())

	user2.addAge()
	fmt.Println(user2.canDrink(), user2.age)
}
