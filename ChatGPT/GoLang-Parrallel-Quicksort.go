package main

import (
	"fmt"
	"math/rand"
	"time"
)

func quicksort(arr []int, done chan bool) {
	if len(arr) <= 1 {
		done <- true
		return
	}

	pivot := arr[0]

	left := []int{}
	right := []int{}

	for _, num := range arr[1:] {
		if num < pivot {
			left = append(left, num)
		} else {
			right = append(right, num)
		}
	}

	leftDone := make(chan bool)
	rightDone := make(chan bool)

	go quicksort(left, leftDone)
	go quicksort(right, rightDone)

	<-leftDone
	<-rightDone

	copy(arr, append(append(left, pivot), right...))

	done <- true
}

func main() {
	rand.Seed(time.Now().UnixNano())
	arr := make([]int, 1000000)
	for i := range arr {
		arr[i] = rand.Intn(1000000)
	}

	done := make(chan bool)
	go quicksort(arr, done)

	<-done

	fmt.Println(arr[:10], arr[len(arr)-10:])
}
