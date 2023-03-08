package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Hello World!")
	fmt.Println("Your random number is", rand.Intn(999))
	fmt.Println("The time is", time.Now())
}
