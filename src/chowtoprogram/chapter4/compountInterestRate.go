package chapter4
import (
	"fmt"
	"math"
)

func main() {
	principal := 1000.0
	rate := 0.05

	var amount float64
	var year int

	fmt.Printf("%4s%21s\n", "Year", "Amount on deposit")
	for year = 1; year < 10; year++ {
		amount = principal * math.Pow(1.0 + rate, float64(year))

		fmt.Printf("%4d%21.2f\n", year, amount)
	}
}