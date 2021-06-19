package caffeine

import "fmt"

type Beverage interface {
	boilWater()
	brew()
	pourInCup()
	addCondiments()
}

type BeverageCaffeine struct {
	Bev Beverage
}

func (b *BeverageCaffeine) boilWater() {
	fmt.Println("Boiling Water...")
}

func (b *BeverageCaffeine) pourInCup() {
	fmt.Println("Pour In Cup...")
}

func (b *BeverageCaffeine) PrepareRecipe() {
	b.boilWater()
	b.Bev.brew()
	b.pourInCup()
	b.Bev.addCondiments()
}

type Coffee struct {
	BeverageCaffeine
}

func (c *Coffee) brew() {
	fmt.Println("Brewing Coffee...")
}

func (c *Coffee) addCondiments() {
	fmt.Println("Adding Sugar and Milk to the Coffee.")
}

type Tea struct {
	BeverageCaffeine
}

func (c *Tea) brew() {
	fmt.Println("Brewing Tea...")
}

func (c *Tea) addCondiments() {
	fmt.Println("Adding Extra lemon to the Tea.")
}
