package chapter3
import "fmt"

func examResults() {
	var result int
	passes := 0
	failures := 0
	students := 1

	for students <= 10 {
		fmt.Println("Enter result(1=pass, 2=fail): ")
		fmt.Scanf("%d", &result)

		if result == 1 {
			passes++
		} else {
			failures++
		}

		students++
	}

	fmt.Printf("Passed: %d\n", passes)
	fmt.Printf("Failed: %d\n", failures)

	if passes > 8 {
		fmt.Println("Raise tuition")
	}
}