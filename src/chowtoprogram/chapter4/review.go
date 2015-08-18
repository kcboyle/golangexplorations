package main
import (
	"fmt"
	"math"
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {
	qFour28()
}

func qFour3() {
	sum := 0
	for i:=1; i <100; i++ {
		sum += i
	}

	fmt.Printf("Sum of nums 1-100 is %d\n\n", sum)

	num := 333.546372
	// "-" will left justify output
	fmt.Printf("Field Width 15, Precision 1: %-15.1f\n", num)
	fmt.Printf("Field Width 15, Precision 2: %-15.2f\n", num)
	fmt.Printf("Field Width 15, Precision 3: %-15.3f\n", num)
	fmt.Printf("Field Width 15, Precision 4: %-15.4f\n", num)
	fmt.Printf("Field Width 15, Precision 5: %-15.5f\n\n", num)

	fmt.Printf("2.5 raised to the power of 3 is %10.2f\n\n", math.Pow(2.5, 3.0))

	for i:=1; i <= 20; i++ {
		fmt.Printf("%d", i)
		if i%5 == 0{
			fmt.Printf("\n")
		} else {
			fmt.Printf("\t")
		}
	}
}

func qFour7() {
	fmt.Printf("1: ")
	for i:=1; i <= 7; i++ {
		fmt.Printf("%d", i)
		if i == 7 {
			continue
		}
		fmt.Printf(",")
	}

	fmt.Printf("\n\n2: ")
	x := 3
	for i:= 0; i < 5; i++ {
		fmt.Printf("%d", x)
		if i == 4 {
			continue
		}
		x += 5
		fmt.Printf(",")
	}

	fmt.Printf("\n\n3: ")
	x = 20
	for i:= 0; i < 6; i++ {
		fmt.Printf("%d", x)
		if i == 5 {
			continue
		}
		x -= 6
		fmt.Printf(",")
	}

	fmt.Printf("\n\n4: ")
	x = 19
	for i:= 0; i < 5; i++ {
		fmt.Printf("%d", x)
		if i == 4 {
			continue
		}
		x += 8
		fmt.Printf(",")
	}
}

func qFour9() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter a number of integers to be summed,")
	fmt.Println("then enter that many integers.")
	intStr, _ := reader.ReadString('\n')
	numInts, _ := strconv.ParseInt(strings.TrimSpace(intStr), 10, 0)

	sum := int64(0)
	for i := int64(0); i < numInts; i++ {
		numStr, _ := reader.ReadString('\n')
		num, _ := strconv.ParseInt(strings.TrimSpace(numStr), 10, 0)
		sum += num
	}

	fmt.Printf("The total sum is %d\n", sum)
}

func qFour10() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter integers and the average will be calculated.")
	fmt.Println("Enter 9999 to end the input.")

	intStr, _ := reader.ReadString('\n')
	numInt, _ := strconv.ParseInt(strings.TrimSpace(intStr), 10, 0)
	sum := int64(0)
	counter := int64(0)
	for numInt != int64(9999) {
		counter++
		sum += numInt
		intStr, _ := reader.ReadString('\n')
		numInt, _ = strconv.ParseInt(strings.TrimSpace(intStr), 10, 0)
	}

	if sum != 0 {
		fmt.Printf("Average is: %.2f", float64(sum)/float64(counter))
	}
}

func qFour11() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter a number of integers to find the minimum of,")
	fmt.Println("then enter that many integers.")

	intStr, _ := reader.ReadString('\n')
	numInts, _ := strconv.ParseInt(strings.TrimSpace(intStr), 10, 0)

	minimum := int64(math.MaxInt64)
	for i := 0; int64(i) < numInts; i++ {
		intStr, _ := reader.ReadString('\n')
		num, _ := strconv.ParseInt(strings.TrimSpace(intStr), 10, 0)

		if minimum > num {
			minimum = num
		}
	}

	if numInts > 0 {
		fmt.Printf("The smallest was %d", minimum)
	}
}

func qFour12() {
	sum := 0

	for i:=2; i <= 30; i++ {
		if i%2 == 0 {
			sum += i
		}
	}

	fmt.Printf("Sum of all even integers from 2 to 30 is %d\n", sum)
}

func qFour13() {
	sum := 0

	for i:= 1; i <= 15; i++ {
		if i%2 != 0 {
			sum += i
		}
	}
	fmt.Printf("Sum of all odd integers from 1 to 15 is %d\n", sum)
}

func qFour14() {
	for i:=1; i <= 5; i++ {
		factorial := i
		for m:=1; m<i; m++ {
			factorial *= i-m
		}
		fmt.Printf("The factorial of %d is %d\n", i, factorial)
	}
}

func qFour16() {
	fmt.Println("(A)")
	for i:=0;i<=10;i++ {
		for j:=1;j<=i;j++ {
			fmt.Printf("*")
			if j == i {
				fmt.Printf("\n")
			}
		}
	}

	fmt.Println("(B)")
	for i:=0;i<=10;i++ {
		for j:=10;j>=i;j-- {
			fmt.Printf("*")
			if j == i {
				fmt.Printf("\n")
			}
		}
	}

	fmt.Println("(C)")
	for i:=0;i<=10;i++ {
		for k:=0;k<=i;k++ {
			fmt.Printf(" ")
		}
		for j:=10;j>=i;j-- {
			fmt.Printf("*")
			if j == i {
				fmt.Printf("\n")
			}
		}
	}

	fmt.Println("(D)")
	for i:=0;i<=10;i++ {
		if i != 0 {
			for k := i; k<=10; k++ {
				fmt.Printf(" ")
			}
		}
		for j:=1;j<=i;j++ {
			fmt.Printf("*")
			if j == i {
				fmt.Printf("\n")
			}
		}
	}
}

func qFour18() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter 5 numbers between 1 and 30")
	fmt.Println("and a pretty histogram will print")

	var nums []int64
	for i:=0;i<5;i++ {
		intStr, _ := reader.ReadString('\n')
		numInt, _ := strconv.ParseInt(strings.TrimSpace(intStr), 10, 0)

		nums = append(nums, numInt)
	}
	for i:=0;i<len(nums);i++ {
		for j:=0; int64(j)<nums[i]; j++ {
			fmt.Printf("*")
		}
		fmt.Printf("\n")
	}
}

func qFour19() {
	reader := bufio.NewReader(os.Stdin)
	productAndRetailTotals := make(map[int]float64)

	fmt.Println("Enter a product number or -1 to end ")
	productStr, _ := reader.ReadString('\n')
	productNum, _ := strconv.ParseInt(strings.TrimSpace(productStr), 10, 0)

	for productNum != -1 {
		fmt.Println("Enter the quantity sold")
		soldStr, _ := reader.ReadString('\n')
		soldNum, _ := strconv.ParseInt(strings.TrimSpace(soldStr), 10, 0)

		switch productNum {
		case 1:
			productAndRetailTotals[1] += float64(soldNum) * 2.98
		case 2:
			productAndRetailTotals[2] += float64(soldNum) * 4.50
		case 3:
			productAndRetailTotals[3] += float64(soldNum) * 9.98
		case 4:
			productAndRetailTotals[4] += float64(soldNum) * 4.49
		case 5:
			productAndRetailTotals[5] += float64(soldNum) * 6.87
		}
		fmt.Println("Enter a product number or -1 to end ")
		productStr, _ := reader.ReadString('\n')
		productNum, _ = strconv.ParseInt(strings.TrimSpace(productStr), 10, 0)
	}
	for i := 1; i<=len(productAndRetailTotals); i++ {
		fmt.Printf("Retail total for product %d is %.2f\n", i, productAndRetailTotals[i])
	}
}

func qFour28() {
	reader := bufio.NewReader(os.Stdin)
	employeeWages := make(map[int][]*float64)

	fmt.Println("Enter paycode number of -1 to end")
	codeStr, _ := reader.ReadString('\n')
	codeNum, _ := strconv.ParseInt(strings.TrimSpace(codeStr), 10, 0)
	for codeNum != -1 {
		switch codeNum {
		case 1:
			fmt.Println("Enter weekly salary")
			workedStr, _ := reader.ReadString('\n')
			workedNum, _ := strconv.ParseFloat(strings.TrimSpace(workedStr), 0)
			employeeWages[1] = append(employeeWages[1],&workedNum)
		case 2:
			fmt.Println("Enter hourly wage")
			wageStr, _ := reader.ReadString('\n')
			wageNum, _ := strconv.ParseFloat(strings.TrimSpace(wageStr), 0)
			fmt.Println("Enter number of hours worked")
			workedStr, _ := reader.ReadString('\n')
			workedNum, _ := strconv.ParseInt(strings.TrimSpace(workedStr), 10, 0)
			var wages float64
			if workedNum > 40 {
				wages = wageNum * float64(40) + 1.5*wageNum*(float64(workedNum-40))
			} else {
				wages = wageNum * 40
			}
			employeeWages[2] = append(employeeWages[2], &wages)
		case 3:
			fmt.Println("Enter gross weekly sales")
			salesStr, _ := reader.ReadString('\n')
			salesNum, _ := strconv.ParseFloat(strings.TrimSpace(salesStr), 0)
			wages := 250.00 + 0.057*salesNum
			employeeWages[3] = append(employeeWages[3], &wages)
		case 4:
			fmt.Println("Enter per product price")
			productStr, _ := reader.ReadString('\n')
			productPrice, _ := strconv.ParseFloat(strings.TrimSpace(productStr), 0)

			fmt.Println("Enter number of products sold")
			producedStr, _ := reader.ReadString('\n')
			producedNum, _ := strconv.ParseInt(strings.TrimSpace(producedStr), 10, 0)
			wages := productPrice * float64(producedNum)

			employeeWages[4] = append(employeeWages[4], &wages)
		}

		fmt.Println("Enter paycode number of -1 to end")
		codeStr, _ := reader.ReadString('\n')
		codeNum, _ = strconv.ParseInt(strings.TrimSpace(codeStr), 10, 0)
	}

	fmt.Println("Wages are:")
	fmt.Println("Employee Code\tWages\n")
	for i:=1; i<=len(employeeWages);i++ {
		for j:=0; j<len(employeeWages[i]);j++ {
			fmt.Printf("%d\t\t$%.2f\n",i, *employeeWages[i][j])
		}
	}
}