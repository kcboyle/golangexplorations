package chapter4
import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

func switchStmt() {
	aCount := 0
	bCount := 0
	cCount := 0
	dCount := 0
	fCount := 0

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter the letter grades")
	fmt.Println("Enter -1 to end input")

	gradeStr, _ := reader.ReadString('\n')
	grade := strings.Replace(gradeStr, "\n", "", -1 )
	
	for grade != "-1" {
		switch grade {
		case "A":
		case "a":
			aCount++

		case "B":
		case "b":
			bCount++

		case "C":
		case "c":
			cCount++

		case "D":
		case "d":
			dCount++

		case "F":
		case "f":
			fCount++

		case "-1":
			break

		default:
			fmt.Println("Incorrect grade entered.")
			fmt.Println("Enter a new grade")
		}
		gradeStr, _ = reader.ReadString('\n')
		grade = strings.Replace(gradeStr, "\n", "", -1 )
	}

	fmt.Println("\nTotals for each letter grade are:")
	fmt.Println("A: ", aCount)
	fmt.Println("B: ", bCount)
	fmt.Println("C: ", cCount)
	fmt.Println("D: ", dCount)
	fmt.Println("F: ", fCount)
}
