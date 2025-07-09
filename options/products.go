package options

type Product struct {
	Name  string
	Price float64
}

type Option func(*Product)

func NewProduct(options ...Option) *Product {
	p := &Product{}
	for _, option := range options {
		option(p)
	}
	return p
}

func WithName(name string) Option {
	return func(p *Product) {
		p.Name = name
	}
}

func WithPrice(price float64) Option {
	return func(p *Product) {
		p.Price = price
	}
}
