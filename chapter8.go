package main

import "fmt"

func main() {

	ch := make(chan bool)

	go func() {
		s := ""
		for i := 0; i < 100; i++ {
			s = s + fmt.Sprint(i, ",")
		}
		fmt.Println(s)
		ch <- true
	}()

	<-ch
}
