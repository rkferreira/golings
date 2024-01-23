// generics1
// Make me compile!

package main

import "fmt"


func main() {
	print("Hello, World!")
	print(42)
}

func print[v string | int](value v) {
	fmt.Println(value)
}
