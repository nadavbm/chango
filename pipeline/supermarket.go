package pipeline

import (
	"fmt"

	"github.com/nadavbm/chango/strategy"
)

type product struct {
	name     string
	price    int
	location string
}

func LuckySupermarket() {
	products := []string{"milk", "bread", "onions", "cucmber", "cookies", "ketchop", "mayo", "meat", "cheese", "salami", "tofu", "apples", "grapes", "chilli", "avocado"}
	locations := []string{"fridge", "upper-shelve", "warehouse", "fruit-corner", "vegtebale-place"}
	bringProductsToStore := importer(products...)
	arrangeProductsInShelves := organazier(bringProductsToStore, locations...)
	setPriceForEachProduct := accountant(arrangeProductsInShelves)
	cashier(setPriceForEachProduct)
}

// importer brings the products to the store
func importer(products ...string) <-chan string {
	out := make(chan string)
	go func() {
		for _, p := range products {
			out <- p
		}
		close(out)
	}()
	return out
}

// organazier arrange the products in the needed location in the store
func organazier(in <-chan string, locations ...string) <-chan product {
	out := make(chan product)
	intGenerator := strategy.Int{}
	go func() {
		for p := range in {
			prdt := product{
				name:     p,
				location: locations[intGenerator.Integer(4, 0)],
			}
			out <- prdt
		}
		close(out)
	}()
	return out
}

// accountant set the price for each product
func accountant(in <-chan product) <-chan product {
	out := make(chan product)
	intGenerator := strategy.Int{}
	go func() {
		for p := range in {
			p.price = intGenerator.Integer(12, 1)
			out <- p
		}
		close(out)
	}()
	return out
}

// cashier print the recipt
func cashier(in <-chan product) {
	count := 0
	for p := range in {
		count++
		fmt.Printf("%d. %s = %d (taken from %s)\n", count, p.name, p.price, p.location)
	}
}
