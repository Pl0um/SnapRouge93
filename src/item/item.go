package item

import (
	"fmt"
)

type Item struct {
	Name         string
	Price        int
}

func (i *Item) ToString() {
	fmt.Printf("Je suis un item qui vaut %d €\n", i.Price)
}
