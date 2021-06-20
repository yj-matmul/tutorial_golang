package main

import "log"

type User struct {
	FirstName string
	LastName  string
}

func main() {
	for i := 0; i < 1; i++ {
		log.Println(i)
	}

	mySlice := []string{"dog", "cat", "horse"}

	// range에서는 index가 자동적으로 return
	for _, x := range mySlice {
		log.Println(x)
	}

	myMap := make(map[string]string)
	myMap["fish"] = "fish"
	myMap["hat"] = "hat"

	// map type 에서는 key, value 가 return
	for i, x := range myMap {
		log.Println(i, x)
	}

	var myUser []User

	u1 := User{
		FirstName: "Yeongjae",
		LastName:  "Yu",
	}
	u2 := User{
		FirstName: "Sam",
	}
	myUser = append(myUser, u1)
	myUser = append(myUser, u2)

	for i, x := range myUser {
		log.Println(i, x.FirstName)
	}

}
