package main

import "log"

type Movement interface {
	movingInfo()
}

type Bird struct {
	name string
}

type Dog struct {
	name string
}

type Food struct {
	name string
}

// () after the function mean that "Bird" has this func has behavior
func (b Bird) movingInfo() {
	log.Printf("%v is flying ...", b.name)
}

// giving movingInfo to dog
func (b Dog) movingInfo() {
	log.Printf("%v is flying ...", b.name)
}

func (b Food) FoodInfo() {
	log.Printf("%v has be eaten  ...", b.name)
}
func main() {

	// we created our structures
	dog := Dog{name: "dog"}
	bird := Bird{name: "bird"}
	food := Food{name: "pizza"}

	// we can manually show movingInfo
	// bird.movingInfo()
	// dog.movingInfo()

	// we created list of our interface with the structures that implement our "movingInfo" then you can show them directly
	//movements := []Movement{dog, bird,food} if we do this
	// we will have error telling us that the 'food' doesn't have 'movingInfo' in his behavior
	movements := []Movement{dog, bird}
	for _, v := range movements {
		v.movingInfo()
	}
	food.FoodInfo()

}
