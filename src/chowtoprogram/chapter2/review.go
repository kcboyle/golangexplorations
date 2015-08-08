package main
import (
	"fmt"
	"math"
	"sort"
)


func main() {
	//Question 2.31 (last in chapter
	nums := []int{0,1,2,3,4,5,6,7,8,9,10}
	fmt.Println("number\tsquare\tcube")
	for i := 0; i < len(nums); i++ {
		square := nums[i] * nums[i]
		cube := square * nums[i]
		fmt.Printf("%d\t%d\t%d\n", nums[i], square, cube)
	}
}

func q1() {
	var c, thisVariable, q76354, number int;

	_ = thisVariable
	_ = q76354
	_ = number

	fmt.Println("Enter an integer: ")
	fmt.Scanf("%d", &c)

	if c != 7 {
		fmt.Printf("The variable %d is not equal to 7\n", c)
	}

	fmt.Println("This is a Go program\n")
	fmt.Println("This is a Go\nprogram\n")
	fmt.Println("This\nis\na\nGo\nprogram\n")
	fmt.Println("This\tis\ta\tGo\tprogram\n")
}

func q2() {
	var x, y, z int

	fmt.Println("This function will calculate the product of 3 integers")
	fmt.Println("Please enter 3 integers: ")
	fmt.Scanf("%d", &x)
	fmt.Scanf("%d", &y)
	fmt.Scanf("%d", &z)

	fmt.Printf("The product is: %d", x * y * z)
}

func qTwo19() {
	var a,b,c int;
	fmt.Println("Input three different integers")
	fmt.Scanf("%d", &a)
	fmt.Scanf("%d", &b)
	fmt.Scanf("%d", &c)

	fmt.Println("Sum is ", a + b + c)
	fmt.Println("Average is ", (a + b + c)/3)
	fmt.Println("Product is ", a * b * c)
	fmt.Println("Smallest is ", smallest(a,b,c))
	fmt.Println("Largest is ", largest(a,b,c))
}

func smallest(a,b,c int) int {
	if a < b && a < c {
		return a
	}
	if b < c && b < a {
		return b
	}
	return c
}

func largest(a,b,c int) int {
	if a > b && a > c {
		return a
	}
	if b > c && b > a {
		return b
	}
	return c
}

func qTwo20() {
	var radius int
	var circumference, area float64
	fmt.Println("Enter the radius of a circle: ")
	fmt.Scanf("%d", &radius)

	circumference = 2.0 * math.Pi * float64(radius)
	area = math.Pi * math.Pow(float64(radius), 2)
	fmt.Printf("The circumference is %v\n", circumference)
	fmt.Printf("The area is %v\n", area)
}

func qTwo21() {
	fmt.Println("*********        ***         *         *")
	fmt.Println("*       *      *     *      ***       * *")
	fmt.Println("*       *     *       *    *****     *   *")
	fmt.Println("*       *     *       *      *      *     *")
	fmt.Println("*       *     *       *      *     *       *")
	fmt.Println("*       *     *       *      *      *     *")
	fmt.Println("*       *     *       *      *       *   *")
	fmt.Println("*       *      *     *       *        * *")
	fmt.Println("*********        ***         *         *")
}

func qTwo23() {
	var int1, int2, int3, int4, int5 int
	//read in 5 ints, determine, and print the largest and smallest
	fmt.Printf("Enter 5 integers, and I will determine the smallest")
	fmt.Scanf("%d", &int1)
	fmt.Scanf("%d", &int2)
	fmt.Scanf("%d", &int3)
	fmt.Scanf("%d", &int4)
	fmt.Scanf("%d", &int5)

	small := smallest(int1,int2,int3)
	small = smallest(small, int4, int5)
	fmt.Println("The smallest integer is: ", small)
}

func qTwo24() {
	var num int
	var even string

	fmt.Println("This function will tell if an integer is odd or even")
	fmt.Println("Enter an integer: ")
	fmt.Scanf("%d", &num)
	if num % 2 == 0 {
		even = "even"
	} else {
		even = "odd"
	}
	fmt.Printf("The number %d is %s", num, even)
}

func qTwo26() {
	var num1, num2 int
	fmt.Println("This program determines if a first integer is a multiple of a second")
	fmt.Println("Enter two integers: ")
	fmt.Scanf("%d", &num1)
	fmt.Scanf("%d", &num2)

	if num2 % num1 == 0 || num1 % num2 == 0 {
		fmt.Println("The numbers are multiples of each other")
		return
	}
	fmt.Println("The numbers are not multiples")
}

func qTwo30() {
	var num int
	fmt.Println("This function will split an integer into individual digits")
	fmt.Println("Enter a 5 digit number")
	fmt.Scanf("%d", &num)

	var digits []int
	for num > 0 {
		digits = append(digits, num%10)
		num = num/10
	}
	sort.Reverse(sort.IntSlice(digits))
	fmt.Println(digits)
}

func qTwo31() {
	nums := []int{0,1,2,3,4,5,6,7,8,9,10}
	fmt.Println("number\tsquare\tcube")
	for i := 0; i < len(nums); i++ {
		square := nums[i] * nums[i]
		cube := square * nums[i]
		fmt.Printf("%d\t%d\t%d\n", nums[i], square, cube)
	}
}