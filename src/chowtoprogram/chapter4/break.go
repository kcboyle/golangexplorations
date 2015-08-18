package chapter4
import "fmt"

func breakFor() {
	var x int
	for x=1; x <= 10; x++ {
		if x == 5 {
			break
		}
		fmt.Printf("%d", x)
	}

	fmt.Printf("\nBroke out of the loop at x == %d", x)
}
