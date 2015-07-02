package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your name: ")
	name, _ := reader.ReadString('\n')
	fmt.Printf("Hello %sWelcome to the forest. Your goal is to appease each of the forest's many denizens through trade.", name)
	fmt.Printf("Your quest is completed once you have pleased the matriarch.")


}
