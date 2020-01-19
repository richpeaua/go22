package main

import "fmt"

func Greeting(prefix string, who ...string) {
	fmt.Printf("%s, %v\n", prefix, who)
}

func main() {
	s := []string{"James", "Jasmine"}
	Greeting("goodbye", s...)
}