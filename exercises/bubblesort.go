package exercices

import (
	"fmt"
)

func GetNumbersListFromInput(inputQuantity int) (inputedNumbers []int, err error) {
	for i := 0; i < inputQuantity; i++ {
		var value int

		fmt.Printf("Type the [%d] integer: ", i)
		_, err := fmt.Scan(&value)
		if err != nil {
			return nil, err
		}

		inputedNumbers = append(inputedNumbers, value)
	}

	return inputedNumbers, nil
}

func Swap(numbers []int, index int) {
	current := numbers[index]
	adjacent := numbers[index+1]

	if current < adjacent {
		return
	}

	numbers[index], numbers[index+1] = adjacent, current
}

func BubbleSort(numbers []int) {
	var totalIterations = len(numbers) - 1

	for i := 0; i < totalIterations; i++ {
		for j := 0; j < totalIterations-i; j++ {
			Swap(numbers, j)
		}
	}
}

func ExecuteBubbleSort() {
	numbersList, err := GetNumbersListFromInput(10)
	if err != nil {
		fmt.Println(err)
		return
	}

	BubbleSort(numbersList)
	fmt.Println(numbersList)
}
