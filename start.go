package main

import (
	"fmt"
	"time"
	"tutorial2/our_packages/add"
)

func main() {
	fmt.Println("Hello Go!")
	fmt.Println("The time is currently, ", time.Now())
	fmt.Println("The result is => ", add.Add(8, 8))
}
