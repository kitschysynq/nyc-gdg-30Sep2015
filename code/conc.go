package main

import "fmt"

func β(s string) {
	fmt.Println(s)
}

func main() {
	β("something")

	go β("concurrent")
	go β("happens here")
}
