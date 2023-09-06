package exercices

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
)

const pieces = 4

func sortPart(part []int, numbersChan chan []int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Sorting ->", part)
	sort.Ints(part)
	numbersChan <- part
}

func splitList(numbers []int, ch chan []int, wg *sync.WaitGroup) {
	result := len(numbers) / pieces

	var part []int
	start := 0
	for i := 0; i < pieces; i++ {
		wg.Add(1)
		current := i + 1
		end := current * result

		if current == 4 {
			part = numbers[start:]
		} else {
			part = numbers[start:end]
		}

		start = end
		go sortPart(part, ch, wg)
	}
}

func parseInput(inputs []string) (integers []int, err error) {
	for _, input := range inputs {
		val, err := strconv.Atoi(input)
		if err != nil {
			return nil, err
		}
		integers = append(integers, val)
	}
	return
}

func InputIntegers() ([]int, error) {
	fmt.Print("Enter integers separated by spaces: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	strs := strings.Split(input, " ")

	if len(strs) <= 1 {
		return nil, fmt.Errorf("nothing to sort")
	}

	integers, err := parseInput(strs)
	if err != nil {
		return nil, err
	}

	return integers, nil
}

func ExecuteGoRoutinesSplitFour() {
	var wg sync.WaitGroup
	var numbersChan = make(chan []int, pieces)

	numbers, err := InputIntegers()
	if err != nil {
		fmt.Println(err)
		return
	}

	splitList(numbers, numbersChan, &wg)
	wg.Wait()
	close(numbersChan)

	var finalList []int
	for part := range numbersChan {
		finalList = append(finalList, part...)
	}

	sort.Ints(finalList)
	fmt.Println("Final list:", finalList)
}
