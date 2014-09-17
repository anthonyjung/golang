package main

import "fmt"

func main() {
	x := 5
	zero(&x)
	fmt.Println(x)
}

func zero(xPtr *int) {
	*xPtr = 0
}
