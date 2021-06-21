package menu

type Item struct {
	Name        string
	Description string
	Price       float32
	Vegetarian  bool
}

type Menu interface {
	Iterator() Iterator
}
