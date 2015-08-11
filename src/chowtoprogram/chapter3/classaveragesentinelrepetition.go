package chapter3

import "fmt"

var grade int

func classAverageSentinelRepetition() {

	total := 0
	counter := 0

	getGrade()

	for grade != -1 {
		total += grade
		counter++
		getGrade()
	}

	if counter != 0 {
		average := float32(total /counter)
		fmt.Printf("Class average is %.2f", average)
		return
	}

	fmt.Println("No grades were entered")
}

func getGrade() {
	fmt.Println("Enter grade, -1 to end: ")
	fmt.Scanf("%d", &grade)
}