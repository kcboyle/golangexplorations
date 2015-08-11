package main
import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {

}

func qThree5() {
	x := 1
	sum := 0

	for x <= 10 {
		sum += x
		x++
	}

	fmt.Printf("The sum is %d\n", sum)
}

func qThree7() {
	var x, y int
	fmt.Println("Enter integer x: ")
	fmt.Scanf("%d", &x)

	fmt.Println("Enter integer y: ")
	fmt.Scanf("%d", &y)

	i := 1
	power := 1

	for i <= y {
		power *= x
		i++
	}

	fmt.Printf("%d raised to the power of %d is %d", x, y, power)
}

func qThree17() {
	var gallons, miles int
	mpgTotal := float32(0)
	fillUps := float32(0)

	fmt.Printf("Enter the gallons used (-1 to end): ")
	fmt.Scanf("%d", &gallons)
	for gallons != -1 {
		fmt.Printf("Enter the miles driven: ")
		fmt.Scanf("%d", &miles)
		mpg := float32(miles)/float32(gallons)
		fillUps++
		mpgTotal += mpg
		fmt.Printf("The miles/gallon for this tank was: %.2f\n", mpg)

		fmt.Printf("Enter the gallons used (-1 to end): ")
		fmt.Scanf("%d", &gallons)
	}

	fmt.Printf("The overall average miles/gallon was %.2f\n", mpgTotal/float32(fillUps))
}

func qThree18() {
	var accountNum int
	var accountBalance, totalCharged, totalCredits, creditLimit float32

	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Enter account number (-1 to end):")
	nums, _ := reader.ReadString('\n')
	num, _ := strconv.ParseInt(strings.TrimSpace(nums), 10, 0)
	accountNum = int(num)

	for accountNum != -1 {
		fmt.Printf("Enter beginning balance: ")
		balance, _ := reader.ReadString('\n')
		bal, _ := strconv.ParseFloat(strings.TrimSpace(balance), 32)
		accountBalance = float32(bal)

		fmt.Printf("Enter total charges:")
		charges, _ := reader.ReadString('\n')
		charge, _ := strconv.ParseFloat(strings.TrimSpace(charges), 32)
		totalCharged = float32(charge)

		fmt.Printf("Enter total credits:")
		credits, _ := reader.ReadString('\n')
		credit, _ := strconv.ParseFloat(strings.TrimSpace(credits), 32)
		totalCredits = float32(credit)

		fmt.Printf("Enter credit limit:")
		limits, _ := reader.ReadString('\n')
		limit, _ := strconv.ParseFloat(strings.TrimSpace(limits), 32)
		creditLimit = float32(limit)

		newBalance := accountBalance + totalCharged - totalCredits

		if newBalance > creditLimit {
			fmt.Printf("Account:\t\t%d\n", accountNum)
			fmt.Printf("Credit Limit:\t\t%f\n", creditLimit)
			fmt.Printf("Balance:\t\t%.2f\n", newBalance)
			fmt.Println("Credit Limit Exceeded")
		}

		fmt.Printf("Enter account number (-1 to end):")
		nums, _ := reader.ReadString('\n')
		num, _ := strconv.ParseInt(strings.TrimSpace(nums), 10, 0)
		accountNum = int(num)
		println(accountNum)
	}

}

func qThree19() {
	basePay := 200
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("Enter sales in dollars (-1 to end): ")
	salesStr, _ := reader.ReadString('\n')
	sales, err := strconv.ParseFloat(strings.TrimSpace(salesStr), 64)
	if err != nil {
		fmt.Println("ERROR: ", err.Error())
		return
	}

	for sales != -1 {
		salary :=  float64(basePay) + (0.09 * sales)
		fmt.Printf("Salary is %.2f\n", salary)

		fmt.Printf("Enter sales in dollars (-1 to end): ")
		salesStr, _ := reader.ReadString('\n')
		sales, _ = strconv.ParseFloat(strings.TrimSpace(salesStr), 64)
	}
}

func qThree20() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("Enter loan principal (-1 to end): ")
	principalStr, _ := reader.ReadString('\n')
	principal, _ := strconv.ParseFloat(strings.TrimSpace(principalStr), 64)

	for principal != -1 {
		fmt.Printf("Enter interest rate: ")
		interestStr, _ := reader.ReadString('\n')
		interestRate, _ := strconv.ParseFloat(strings.TrimSpace(interestStr), 64)

		fmt.Printf("Enter term of the loan in days: ")
		termDaysStr, _ := reader.ReadString('\n')
		termDays, _ := strconv.ParseFloat(strings.TrimSpace(termDaysStr), 64)

		charge := principal * interestRate * termDays / 365

		fmt.Printf("The interest charge is $%.2f\n", charge)

		fmt.Printf("Enter loan principal (-1 to end): ")
		principalStr, _ := reader.ReadString('\n')
		principal, _ = strconv.ParseFloat(strings.TrimSpace(principalStr), 64)
	}
}

func qThree21() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("Enter # of hours worked (-1 to end): ")
	hoursStr, _ := reader.ReadString('\n')
	hours, _ := strconv.ParseInt(strings.TrimSpace(hoursStr), 10, 0)

	for hours != -1 {
		salary := 0.00

		fmt.Printf("Enter hourly rate of the worker ($00.00): ")
		rateStr, _ := reader.ReadString('\n')
		rate, _ := strconv.ParseFloat(strings.TrimSpace(rateStr), 64)

		if hours > 40 {
			overtime := hours - 40
			salary = (40.00 * rate) + (float64(overtime) * 1.50 * rate)
		} else {
			salary = 40.00 * rate
		}

		fmt.Printf("Salary is %.2f\n\n", salary)

		fmt.Printf("Enter # of hours worked (-1 to end): ")
		hoursStr, _ := reader.ReadString('\n')
		hours, _ = strconv.ParseInt(strings.TrimSpace(hoursStr), 10, 0)
	}
}

func qThree25() {
	fmt.Println("N\t10*N\t100*N\t1000*N")
	for i := 1; i <= 10; i++ {
		fmt.Printf("%d\t%d\t%d\t%d\n", i, i * 10, i * 100, i * 1000)
	}
}

func qThree33() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter a length of a square: ")
	lengthStr, _ := reader.ReadString('\n')
	length, _ := strconv.ParseInt(strings.TrimSpace(lengthStr), 10, 64)
	for i := int64(0); i < length; i++ {
		for i := int64(0); i < length; i++ {
			fmt.Printf("*")
		}
		fmt.Printf("\n")
	}
}

func qThree34() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter a length of a square: ")
	lengthStr, _ := reader.ReadString('\n')
	length, _ := strconv.ParseInt(strings.TrimSpace(lengthStr), 10, 64)
	for i := int64(0); i < length; i++ {
		if i == 0 || i+1 == length {
			for i := int64(0); i < length; i++ {
				fmt.Printf("*")
			}
			fmt.Printf("\n")
		} else {
			for i := int64(0); i < length; i++ {
				if i == 0 || i+1 == length {
					fmt.Printf("*")
				} else {
					fmt.Printf(" ")
				}
			}
			fmt.Printf("\n")
		}
	}
}

func qThree35() {
	reader := bufio.NewReader(os.Stdin)
	var isPalindrome bool

	fmt.Println("Enter a number: ")
	palindromeStr, _ := reader.ReadString('\n')
	palindrome, _ := strconv.ParseInt(strings.TrimSpace(palindromeStr), 10, 64)

	var digits []int64
	for palindrome>0 {
		digits = append(digits, palindrome%10)
		palindrome /= 10
	}

	length := len(digits)
	for i := 0; i < (length/2); i++ {
		if digits[i] != digits[length-1-i] {
			isPalindrome = false
			break
		}
		isPalindrome = true
	}

	if isPalindrome {
		fmt.Println("It's a palindrome!")
		return
	}
	fmt.Println("It's not a palindrome...")
}

func qThree36() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("Enter a binary number: ")
	binaryStr, _ := reader.ReadString('\n')
	binary, _ := strconv.ParseInt(strings.TrimSpace(binaryStr), 10, 64)

	var digits []int64
	for binary > 0 {
		digits = append(digits, binary%10)
		binary /= 10
	}

	total := 0
	operator := 0
	for i := 0; i < len(digits); i++ {
		if i == 0 {
			operator++
		} else {
			operator *= 2
		}
		if digits[i] == 1 {
			total += operator
		}
	}

	fmt.Printf("The number in decimal form is: %d", total)
}

func qThree38() {
	for i := 0; i < 100; i++ {
		if i%10 == 0 {
			fmt.Printf("\n")
		}
		fmt.Printf("*")
	}
}