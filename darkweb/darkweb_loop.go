package darkweb

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Target struct {
	Email     string
	FirstName string
	LastName  string
	Position  string
}

type Response struct {
	Name    string
	Targets []Target
}

func main() {
	// Make a slice with a cap of 15
	var targets = make([]Target, 0, 15000000)

	// Dump our data into the slice
	for i := 0; i < 15000000; i++ {
		targets = append(targets, Target{
			"chris" + strconv.Itoa(i) + "@mdingrub.com",
			"c",
			"z",
			"human",
		})
	}

	response := Response{
		"Group 123",
		targets,
	}

	targetsJson, err := json.MarshalIndent(response, "", "    ")
	if err != nil {
		log.Fatal("Cannot encode to JSON ", err)
	}

	file, err := os.Create("post-150-mil.json")
	if err != nil {
		fmt.Println(err)
		return
	}

	l, err := file.Write(targetsJson)
	if err != nil {
		fmt.Println(err)
		file.Close()
		return
	}
	fmt.Println(l, "bytes written successfully")
	err = file.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

}
