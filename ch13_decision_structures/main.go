package main

import "log"

func main() {
	var isTrue bool

	isTrue = true

	if isTrue {
		log.Println("isTrue is", isTrue)
	} else {
		log.Println("isTrue is", isTrue)
	}

	myNum := 100
	isTrue = false

	if myNum > 99 && isTrue {
		log.Println("myNum is greater than 99 and isTrue is set to true")
	} else if myNum < 100 && isTrue {
		log.Println("myNum is less than 100 and isTrue is set to True")
	} else if myNum == 101 || isTrue {
		log.Println("2")
	} else if myNum > 1000 && isTrue == false {
		log.Println("3")
	}

	myVar := "cat"

	switch myVar {
	case "cat":
		log.Println("cat is set to cat")
	case "dog":
		log.Println("cat is set to dog")
	case "fish":
		log.Println("cat is set to fish")
	default:
		log.Println("cat is set to something")
		// switch 문은 어느 케이스라도 마주치면 자동적으로 break 이 실행
	}
}
