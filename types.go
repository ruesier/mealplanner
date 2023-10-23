package mealplanner

type Nutrition uint16

const (
	STARCH Nutrition = 1 << iota
	PROTEIN
	VEGETABLE
	FRUIT
)

type Ingredient struct {
	Name   string
	Amount float64
	Unit   string
}

type Meal struct {
	Name        string
	Nutrition   Nutrition
	Ingredients []Ingredient
}
