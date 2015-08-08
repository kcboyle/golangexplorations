package chapter2
import "fmt"

func divideByZero() {
	var a int
	fmt.Println("Enter an integer")
	fmt.Scanf("%d", &a)

	fmt.Println("Result: ", 10/a)

}