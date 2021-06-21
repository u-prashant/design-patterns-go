package menu

func NewDinerMenu() *DinerMenu {
	m := DinerMenu{}
	m.AddItem(Item{
		Name:        "Paneer Do Pyaaza",
		Description: "Vegetable with paneer and pyaaz",
		Price:       99.9,
		Vegetarian:  true,
	})
	m.AddItem(Item{
		Name:        "Dum Aloo",
		Description: "Vegetable with potato",
		Price:       19.9,
		Vegetarian:  true,
	})
	m.AddItem(Item{
		Name:        "Rice",
		Description: "Jeera Rice",
		Price:       39,
		Vegetarian:  true,
	})
	return &m
}

type DinerMenu struct {
	MenuItems [3]Item
	items     int
}

func (d *DinerMenu) Iterator() Iterator {
	return &DinerMenuIterator{MenuItems: d.MenuItems[:]}
}

func (d *DinerMenu) AddItem(item Item) {
	d.MenuItems[d.items] = item
	d.items++
}

type DinerMenuIterator struct {
	MenuItems []Item
	position  int
}

func (d *DinerMenuIterator) HasNext() bool {
	if d.position >= len(d.MenuItems) {
		return false
	}
	return true
}

func (d *DinerMenuIterator) Next() Item {
	d.position++
	return d.MenuItems[d.position-1]
}
