package main

import "design-patterns-go/10_template_pattern/caffeine"

func main() {
	c := caffeine.BeverageCaffeine{}

	tea := caffeine.Tea{}
	c.Bev = &tea
	c.PrepareRecipe()

	coffee := caffeine.Coffee{}
	c.Bev = &coffee
	c.PrepareRecipe()
}
