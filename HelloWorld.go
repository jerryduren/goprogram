package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hlello World!")

	const day = 24 * time.Hour
	fmt.Println(day.Hours())
	fmt.Printf("%v:%v:%v:%v", day.Hours(), day.Minutes(), day.Seconds(), day.Nanoseconds())
}
