package main

import (
	"fmt"
)

func pockerLadder(numbers []int,  channel chan bool) {
	isValid:= true
	isValid = validate(numbers)
	if !isValid {
		channel <- isValid
		close(channel)
		return
	}
	minNumber, maxNumber:= getMinAndMaxNumber(numbers)
	sum:= 0
	index:= 1
	for i, _ := range numbers {
		if i == 0 {
			sum = minNumber
		} else {
			sum = sum + 1
		}
		if index == len(numbers) && sum != maxNumber  {
			isValid = false
			break
		}
		index= index + 1 
	}
	channel <- isValid
	close(channel)
}

func validatePockerHands(numbers []int) bool {
	channel := make(chan bool) 
  go	pockerLadder(numbers, channel)
	return <-channel 
}


func main () {
	numbers2:=[]int{14, 5, 4 ,7, 14}
	response:= validatePockerHands(numbers2)
  fmt.Println("response:",response)
}