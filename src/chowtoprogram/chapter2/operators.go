package main
import "fmt"

func operators() {
	fmt.Println("Enter two integers, and I will tell you")
	fmt.Println("the relationship they satisfy")

	var int1, int2 int
	fmt.Scanf("%d", &int1)
	fmt.Scanf("%d", &int2)

	if int1 == int2 {
		fmt.Printf("%d is equal to %d\n", int1, int2)
	}

	if int1 > int2 {
		fmt.Printf("%d is greater than %d\n", int1, int2)
	}

	if int1 < int2 {
		fmt.Printf("%d is less than %d\n", int1, int2)
	}

	if int1 >= int2 {
		fmt.Printf("%d is greater than or equal to %d\n", int1, int2)
	}

	if int1 <= int2 {
		fmt.Printf("%d is less than or equal to %d\n", int1, int2)
	}
}
