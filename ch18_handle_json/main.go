package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	HairColer string `json:"hair_color"`
	HasDog    bool   `json:"has_dog"`
}

func main() {
	myJson := `
[
	{
		"first_name": "Clark",
		"last_name": "Kent",
		"hair_color": "black",
		"has_dog": true
	},
	{
		"first_name": "Bruce",
		"last_name": "Wayne",
		"hair_color": "yellow",
		"has_dog": false
	}
]`
	// read struct from a json
	var unmarshalled []Person

	err := json.Unmarshal([]byte(myJson), &unmarshalled)
	if err != nil {
		log.Println("Error unmarshalling json", err)
	}

	log.Printf("unmarshalled: %v", unmarshalled)
	for i, x := range unmarshalled {
		log.Printf("for loop %d %v", i, x)
	}

	// wirte json from a struct
	var mySlice []Person

	m1 := Person{
		FirstName: "Wally",
		LastName:  "West",
		HairColer: "red",
		HasDog:    false,
	}

	mySlice = append(mySlice, m1)

	newJson, err := json.MarshalIndent(mySlice, "", "  ")
	if err != nil {
		log.Println("error marshalling", err)
	}

	fmt.Println(string(newJson))

}
