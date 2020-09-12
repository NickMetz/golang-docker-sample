package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("This is a golang docker sample!")
	fmt.Printf("Magic Number of the day is %v \n", RandomNumber())
}

// RandomNumber returns a random number sourced by time
func RandomNumber() int {
	source := rand.NewSource(time.Now().UnixNano())
	gen := rand.New(source)

	return gen.Int()
}
