package main

import (
	"log"

	"github.com/yj-matmul/tutorial_go/helpers"
)

const numPool = 10

func CalculateValue(intChan chan int) {
	randomNumber := helpers.RandomNumber(numPool)
	intChan <- randomNumber
}

func main() {
	intChan := make(chan int)
	defer close(intChan) // defer : main 함수 종료되기 직전에 해당 명령줄이 실행되도록 하는 기능

	go CalculateValue(intChan)

	num := <-intChan
	log.Println(num)
}
