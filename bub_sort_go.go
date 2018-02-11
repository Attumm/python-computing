package main

import "fmt"
import "math/rand"
import "time"

const size int = 5000
const jobs int = 3

var msg = [500]string{"one", "two", "three", "four", "five"}

func main() {
	messages := make(chan int)
	for i := 0; i < jobs; i++ {
		go run(msg[i], messages)
	}

	for i := 0; i < jobs; i++ {
		n := <-messages
		if false {
			fmt.Println(n)
		}
	}
}

func run(name string, messages chan int) {
	sorted := bub_sort(arr_ran())
	fmt.Println("done", name, len(sorted))
	messages <- 1
}

func arr_ran() [size]int {
	rand.Seed(time.Now().UnixNano())
	n := [size]int{}
	for i := 0; i < size; i++ {
		n[i] = rand.Intn(10)
	}
	return n
}

func bub_sort(n [size]int) [size]int {
	for i := 0; i < (size - 1); i++ {
		for j := 0; j < ((size - 1) - i); j++ {
			if n[j] > n[j+1] {
				n[j], n[j+1] = n[j+1], n[j]
			}
		}
	}
	return n
}
