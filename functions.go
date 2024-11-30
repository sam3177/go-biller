package main

import (
	"fmt"
	"strings"
)

func getNameInitials(firstName string, lastName string) (string, string) {

	return strings.ToUpper(string(firstName[0])), strings.ToUpper(string(lastName[0]))
}

func clojure() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}

func greet(name string) {
	fmt.Println("Hello", name)
}

func multiplyBy(multiplier int) func(number float64) float64 {
	return func(number float64) float64 {
		return number * float64(multiplier)
	}
}

func mapSlice(slice []float64, mappingFn func(number float64) float64) []float64 {
	newSlice := []float64{}
	for _, num := range slice {
		newSlice = append(newSlice, mappingFn(num))
	}

	return newSlice
}

// in1, in2 := getNameInitials("john", "Doe")

// println(in1, in2)

// nextInt := clojure()
// println(nextInt())
// println(nextInt())
// println(nextInt())

// newInts := clojure()
// println(newInts())
// println(newInts())
// println(newInts())

var nums []float64 = []float64{1, 2, 3, 4, 5}
