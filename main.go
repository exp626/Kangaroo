package main

import (
	"fmt"
	kg "github.com/exp626/Kangaroo/Kangaroo"
)

var x1 int
var v1 int
var x2 int
var v2 int

func main() {
	fmt.Println("Please enter the data separated by spaces, format: x1 v1 x2 v2")
	fmt.Println("Where x1 v1 x2 v2 integers.")
	_, err := fmt.Scanf("%d %d %d %d", &x1, &v1, &x2, &v2)

	if err != nil {
		fmt.Println("Invalid enter data.")
		fmt.Println("Please enter the data separated by spaces, format: x1 v1 x2 v2")
		fmt.Println("Next time.")
		return
	}

	if kg.KangarooCheck(x1, v1, x2, v2) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
