package menu

func NewPancakeHouseMenu() *PancakeHouseMenu {
	p := PancakeHouseMenu{}
	p.AddItem(Item{
		Name:        "Cheese Burger",
		Description: "Light Cheese Burger",
		Price:       20,
		Vegetarian:  false,
	})
	p.AddItem(Item{
		Name:        "Chicken Bucket",
		Description: "Bucket full of chicken wings",
		Price:       120,
		Vegetarian:  false,
	})
	return &p
}

type PancakeHouseMenu struct {
	MenuItems []Item
}

func (p *PancakeHouseMenu) AddItem(item Item) {
	p.MenuItems = append(p.MenuItems, item)
}

func (p *PancakeHouseMenu) Iterator() Iterator {
	return &PancakeHouseIterator{MenuItems: p.MenuItems}
}

type PancakeHouseIterator struct {
	MenuItems []Item
	position  int
}

func (p *PancakeHouseIterator) HasNext() bool {
	if p.position >= len(p.MenuItems) {
		return false
	}
	return true
}

func (p *PancakeHouseIterator) Next() Item {
	p.position++
	return p.MenuItems[p.position-1]
}
