package main

import "fmt"

func main() {
	basic()

	// adding elements and ranging
	workingWithMaps()

	// Random Notes
	// maps use keys for identifying element values.
	// arrays and slices use index for identifying element values.
}

func basic() {
	// Maps are key value store.
	// Very fast look ups
	// unordered list.

	// map[KEY_TYPE]VALUES_TYPE
	// trailing comma is required.
	m := map[string]int{
		"James": 21,
		"Jim":   18,
	}
	fmt.Println(m)

	fmt.Println(m["James"])

	// If a key isn't found, the map returns a 0
	fmt.Println(m["Gerrad"])

	// comma ok idiom
	// value is returned first, ok is true if the the key exists in the map, and false if it does not.
	v, ok := m["Gerrad"]
	fmt.Println(v)
	fmt.Println(ok)

	// Common shorthand for "IF this value exist"
	if v, ok := m["Gerrad"]; ok {
		fmt.Println("This is the if print", v)
	}
}

func workingWithMaps() {
	m := map[string]int{
		"James": 21,
		"Jim":   18,
	}

	// Adding element to a map.
	m["todd"] = 34

	for k, val := range m {
		fmt.Println(k, val)
	}

	// deleting element from map.
	delete(m, "todd")

	for k, val := range m {
		fmt.Println(k, val)
	}

	// No errors are thrown if trying to delete something that doesn't exist.
	delete(m, "Topher")
}
