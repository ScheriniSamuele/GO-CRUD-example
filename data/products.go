package data

type Product struct {
	PID      int     `json:"id"`
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Quantity int     `json:"quantity"`
}

func InitProducts() []Product {
	products := make([]Product, 0)

	products = append(products, Product{PID: 1, Name: "apples 🍎", Price: 1.09, Quantity: 4})
	products = append(products, Product{PID: 2, Name: "peaches 🍑", Price: 1.45, Quantity: 2})
	products = append(products, Product{PID: 3, Name: "kiwi 🥝", Price: 1.12, Quantity: 0})
	products = append(products, Product{PID: 4, Name: "grapes 🍇", Price: 0.99, Quantity: 8})

	return products
}