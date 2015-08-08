package chapter2

import (
	"fmt"
)


func main() {
	var int1, int2 int

	fmt.Println("Enter the first integer:")
	fmt.Scanf("%d", &int1)

	fmt.Println("Enter second integer:")
	fmt.Scanf("%d", &int2)

	fmt.Printf("Sum is %d", int1 + int2)
}