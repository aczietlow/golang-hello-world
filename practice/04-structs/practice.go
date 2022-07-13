package main

import (
	"fmt"
)

func main() {
	// Create your own type person which will have an underlying type of "struct" so that it can store the following data:
	// first name
	// last name
	// favorite ice cream flavors
	// Create two VALUES of TYPE person. Print out the values ranging over the elements in the slice which stores the favorite flavors
	one()

	// Take the code from the one() and store the values of the type person in a map with the key of the last name.
	// Access each value in the map. Print out the values, ranging over the slice.
	//two()

	// Create a new type: vehicle
	// - underlying type is a struct
	// -- the fields are 'doors' and 'color'
	// Create two new types: truck and sedan.
	// - the underlying type of each is a struct
	// - Embed the vehicle type for both struct and sedan
	// - Give truck the field 'fourWheel' which will be set to bool
	// - Give sedan the field 'luxury' which will be set to bool.
	// Using the vehicle, truck, and sedan structs:
	// - Using a composite literal, create a value of type truck and assign values to the fields;
	// - using a composite literal, create a value of type sedan and assign values to the fields
	// print out each of these values
	// print out a single field from each of these values
	three()

	// Create a use an anonymous struct.
	four()
}

func one() {
	type person struct {
		first   string
		last    string
		flavors []string
	}
	p1 := person{
		first:   "Jim",
		last:    "Dandy",
		flavors: []string{"chocolate", "vanilla"},
	}

	fmt.Println(p1)
	for _, flavor := range p1.flavors {
		fmt.Println(flavor)
	}

	p2 := person{
		first:   "alaina",
		last:    "Zietlow",
		flavors: []string{"vanilla", "mouse tracks", "almond milk"},
	}

	// Maps are good for key value lookup.
	m := map[string]person{
		p1.last: p1,
		p2.last: p2,
	}

	for _, v := range m {
		fmt.Println(v.first)
		fmt.Println(v.last)
		for i, val := range v.flavors {
			fmt.Println(i, val)
		}
		fmt.Println("----------")
	}
}

func three() {

	type vehicle struct {
		doors int
		color string
	}

	type truck struct {
		vehicle
		fourWheel bool
	}

	type sedan struct {
		vehicle
		luxury bool
	}

	t := truck{
		vehicle: vehicle{
			doors: 4,
			color: "black",
		},
		fourWheel: true,
	}

	fmt.Println(t)

	s := sedan{
		vehicle: vehicle{
			doors: 2,
			color: "pink",
		},
		luxury: false,
	}

	fmt.Println(s)
	// Can access inner type from the outer type.
	fmt.Println(s.doors)
}

func four() {
	// I got carried away a bit.
	type totalComp func(int, int) int

	type user struct {
		username     string
		password     string
		salary       totalComp
		vacationDays func(int) int
	}

	u1 := user{
		username: "bob",
		password: "12345",
		salary: func(salary int, quarterlyVariableComp int) int {
			return salary * (quarterlyVariableComp * 4)
		},
		vacationDays: func(tenure int) int {
			var days int
			switch {
			case tenure < 2:
				days = 10
			case tenure < 5:
				days = 15
			case tenure >= 5:
				days = 20
			}
			return days
		},
	}
	fmt.Println(u1)

	org := struct {
		name, industry string
		team           map[string]int
		favDrinks      []string
	}{
		name:     "redhat",
		industry: "hats",
		team: map[string]int{
			"Reese":  1111,
			"Ryan":   2222,
			"Kaylan": 3333,
		},
		favDrinks: []string{
			"whiskey drink",
			"vodka drink",
			"lager drink",
		},
	}

	fmt.Println(org.name)
	fmt.Println(org.industry)

	for k, v := range org.team {
		fmt.Println(k, v)
	}

	for _, v := range org.favDrinks {
		fmt.Println(v)
	}

}
