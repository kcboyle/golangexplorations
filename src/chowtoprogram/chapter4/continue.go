package chapter4
import "fmt"

func continueFor() {
	var x int
	for x=1; x<=10; x++ {
		if x == 5 {
			continue
		}

		fmt.Printf("%d ", x)
	}
	fmt.Printf("\nUsed continue to skip printing the value 5\n")
}
