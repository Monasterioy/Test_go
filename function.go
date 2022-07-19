package main


func getMinAndMaxNumber(numbers []int) (int,int) {
	minNumber:= 0
	maxNumber:= 0
	if  numbers[0] != 14 {
		minNumber = numbers[0]
		maxNumber = numbers[0]
	}  else {
		minNumber= 1
		maxNumber= 1
	}

	for i, number := range numbers {
		if number < minNumber && i != 0 && number != 14  {
			minNumber = number
		} else if number < minNumber && i != 0 && number == 14 {
			minNumber = 14
		}

		if number > maxNumber && i != 0 && number != 14  {
			maxNumber = number
		} else if number > maxNumber && i != 0 && number == 14 {
			maxNumber = 14
		}
	}


  return minNumber, maxNumber
}

func validate(numbers []int) bool {
	isValid := true
	if len(numbers) > 7 || len(numbers) < 2 || len(numbers) == 0{
		isValid = false
	}
	for _, number := range numbers {
		if number < 2 || number > 14 {
			isValid = false
			break
		}
	}
	return isValid
}
