// +build ignore

package main

import "fmt"

func main() {
	v := 42
	fmt.Printf("v is of type %T\n", v)
	fmt.Printf("v is of type %T\n", 3.14)
	fmt.Printf("v is of type %T\n", 0.867 + 0.5i)
}