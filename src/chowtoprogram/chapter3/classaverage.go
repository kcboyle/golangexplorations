package chapter3
import "fmt"


func classAverage() {
	var grade int

	total := 0
	counter := 1

	for counter <= 10 {
		fmt.Println("Enter grade: ")
		fmt.Scanf("%d", &grade)

		total += grade
		counter++
	}

	average := total / (counter - 1)
	fmt.Printf("Class average is %d\n", average)
}