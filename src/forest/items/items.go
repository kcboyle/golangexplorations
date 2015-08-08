package items

import (
	"bufio"
	"log"
	"os"
)

type item struct {
	allItems      []string
	tradableItems chan string
	wishableItems chan string
}

func NewItems() *item {
	return &item{
		allItems:      populateAllItems(),
		tradableItems: make(chan string),
		wishableItems: make(chan string),
	}
}

func populateAllItems() []string {
	var items []string
	itemFile := "itemList.txt"
	file, err := os.Open(itemFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		items = append(items, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return items
}

func populateTradableItems() {

}

func (i *item) GetAllItems() []string {
	return i.allItems
}
