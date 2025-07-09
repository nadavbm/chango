package workerpool

import (
	"fmt"
	"time"

	"github.com/nadavbm/chango/decorator"
	"github.com/nadavbm/chango/singleton"
	"github.com/nadavbm/chango/strategy"
)

const numberOfProductsToArrange = 20

type product struct {
	name  string
	price int
}

func WorkInSupermarket(logger decorator.Logger, cfg *singleton.Config) {
	jobs := make(chan int, numberOfProductsToArrange)
	results := make(chan product, numberOfProductsToArrange)

	for w := 1; w <= cfg.Workers; w++ {
		go work(logger, w, jobs, results)
	}

	// Send jobs to the channel
	for j := 1; j <= numberOfProductsToArrange; j++ {
		jobs <- j
	}
	close(jobs)

	time.Sleep(3 * time.Second)
	// Show new products on the shelves
	for a := 1; a <= numberOfProductsToArrange; a++ {
		fmt.Printf("%v\n", <-results)
	}
}

func work(logger decorator.Logger, id int, jobs <-chan int, results chan<- product) {
	time.Sleep(1 * time.Second)
	productsPriceGenerator := strategy.Int{}
	productsNameGenerator := strategy.Str{}
	for j := range jobs {
		product := product{
			name:  fmt.Sprintf("%d. %s", j, productsNameGenerator.String(productsPriceGenerator.Integer(12, 4))),
			price: productsPriceGenerator.Integer(10, 2),
		}
		productName := productsNameGenerator.String(productsPriceGenerator.Integer(12, 4))
		logger.Info(fmt.Sprintf("worker %d stick a new product price to product %s", id, productName))
		results <- product
	}
}
