package main

import (
	"fmt"
	"strings"

	"github.com/ldej/advent-of-code/tools"
	"github.com/ldej/advent-of-code/tools/sets"
)

func main() {
	fmt.Println("Part 1")

	result := run("./2020/day21/example1.txt")
	fmt.Println("Example:", result)

	result = run("./2020/day21/input.txt")
	fmt.Println("Result:", result)
}

type Product struct {
	Ingredients *sets.StringSet
	Allergens   []string
}

func run(input string) int {
	lines := tools.ReadStrings(input)

	var allergens = sets.NewStringSet()

	var products []Product
	for _, line := range lines {
		items := strings.Split(line, " (contains ")
		_ingredients := strings.Split(items[0], " ")
		_allergens := strings.Split(strings.TrimRight(items[1], ")"), ", ")

		for _, allergen := range _allergens {
			allergens.Add(allergen)
		}

		products = append(products, Product{
			Ingredients: sets.NewStringSet(_ingredients...),
			Allergens:   _allergens,
		})
	}

	var possibleAllergensMap = make(map[string]*sets.StringSet)

	for _, product := range products {
		for _, allergen := range product.Allergens {
			if _, found := possibleAllergensMap[allergen]; found {
				possibleAllergensMap[allergen] = possibleAllergensMap[allergen].Intersect(product.Ingredients)
			} else {
				possibleAllergensMap[allergen] = product.Ingredients
			}
		}
	}

	var definiteAllergensMap = make(map[string]string)

	for len(possibleAllergensMap) > 0 {
		for allergen, ingredients := range possibleAllergensMap {
			if ingredients.Len() == 1 {
				ingredient := ingredients.Items()[0]
				definiteAllergensMap[allergen] = ingredient

				for allergen1, ingredients1 := range possibleAllergensMap {
					possibleAllergensMap[allergen1] = ingredients1.Remove(ingredient)
				}
				delete(possibleAllergensMap, allergen)
				break
			}
		}
	}

	var result int
	for _, product := range products {
		for _, ing := range product.Ingredients.Items() {
			var found bool
			for _, allergen := range definiteAllergensMap {
				if ing == allergen {
					found = true
				}
			}
			if !found {
				result++
			}
		}
	}

	return result
}
