package main

import (
	"fmt"
	"math"
	"time"
)

type Function func(float64) float64

func ProfileDecorator(fn Function) Function {
	return func(params float64) float64 {
		start := time.Now()
		result := fn(params)
		elaplsed := time.Now().Sub(start)
		fmt.Println("Function completed with time: ", elaplsed)
		return result
	}
}

func SquareRoot(n float64) float64 {
	return math.Sqrt(n)
}

func main() {
	decoratedSquareRoot := ProfileDecorator(SquareRoot)
	fmt.Println(decoratedSquareRoot(16))
}
