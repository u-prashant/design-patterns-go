package menu

import "fmt"

type Waitress struct {
	PancakeHouse Menu
	DinerMenu    Menu
}

func (w *Waitress) PrintMenu() {
	pancakeIterator := w.PancakeHouse.Iterator()
	dinerIterator := w.DinerMenu.Iterator()
	fmt.Println("BREAKFAST")
	w.print(pancakeIterator)
	fmt.Println("LUNCH")
	w.print(dinerIterator)
}

func (w *Waitress) print(iterator Iterator) {
	for iterator.HasNext() {
		item := iterator.Next()
		fmt.Printf("Name: %s, Description: %s, Price: %f\n", item.Name, item.Description, item.Price)
	}
}
