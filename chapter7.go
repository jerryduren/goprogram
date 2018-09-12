package main

import (
	"flag"
	"fmt"
	"time"
)

// Demo command line parameter: chapter7 -period 10s -U neil -P ddd
var period = flag.Duration("period", 1*time.Second, "sleep period") //usage: 参数的解释
var turnoff = flag.Bool("switch", false, "switch")
var name = flag.String("U", "Jerry", "Input Login Name")
var password = flag.String("P", "197419802008", "Input password")

func main() {
	var any interface{}

	any = map[string]int{"one": 1, "two": 2, "three": 3}
	switch value := any.(type) {
	case map[string]int:
		for k, v := range value {
			fmt.Println("key=", k, "; value=", v)
		}
	}
	any = "hello"
	fmt.Println(any)

	flag.Parse()
	fmt.Println("Now it is time", time.Now(), ", I will sleep for ", *period)
	time.AfterFunc(*period, func() {
		fmt.Println("Now it is time", time.Now(), ", I has sleeped ", *period)
	})

	time.Sleep(*period + time.Second)

	fmt.Println("Switch is ", *turnoff)

	fmt.Println("Login Name:", *name, "Login Password:", *password)
}
