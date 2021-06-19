package main

import (
	"log"
	"sort"
)

type User struct {
	FirstName string
	LastName  string
}

func main() {
	// map
	myMap := make(map[string]string)

	myMap["dog"] = "Mango"
	myMap["other-dog"] = "Cola"

	log.Println(myMap["dog"])
	log.Println(myMap["other-dog"])

	myMap2 := make(map[string]int)

	myMap2["First"] = 1
	myMap2["First"] = 0

	log.Println(myMap2["First"])

	// can have struct
	myMap3 := make(map[string]User)

	me := User{
		FirstName: "Yeongjae",
		LastName:  "Yu",
	}

	myMap3["me"] = me

	log.Println(myMap3["me"])

	// slice
	var mySlice []int

	mySlice = append(mySlice, 2)
	mySlice = append(mySlice, 1)
	mySlice = append(mySlice, 3)

	log.Println(mySlice)
	sort.Ints(mySlice)
	log.Println(mySlice)

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	log.Println(numbers)
	log.Println(numbers[6:9])

}
