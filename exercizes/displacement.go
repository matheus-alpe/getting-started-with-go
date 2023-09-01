package exercices

import (
	"fmt"
	"log"
)

type fnTimeDisplacement func(float64) float64

func GenDisplaceFn(a, vo, so float64) fnTimeDisplacement {

	return func(time float64) float64 {
		return ((1.0 / 2.0) * a * time * time) + (vo * time) + so
	}
}

func InputDisplacementParams() (params []float64, err error) {
	messages := [4]string{
		"acceleration",
		"initial velocity",
		"initial displacement",
		"time",
	}

	fmt.Println("Initializing displacement formula")
	for _, message := range messages {
		fmt.Printf("Input %s: ", message)
		var input float64
		_, err := fmt.Scan(&input)
		if err != nil {
			return nil, err
		}

		params = append(params, input)
	}

	return params, nil
}

func ExecuteDisplacement() {
	params, err := InputDisplacementParams()
	if err != nil {
		log.Fatal(err)
	}
	acceleration := params[0]
	initialVelocity := params[1]
	initialDisplacement := params[2]
	time := params[3]

	fn := GenDisplaceFn(acceleration, initialVelocity, initialDisplacement)

	fmt.Println(fn(time))
}
