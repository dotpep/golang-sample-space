package main

import "fmt"

// Composition e.g. 2

type Product struct {
	Name  string
	Price int
}

type CartShopper interface {
	AddProduct()
	GetTotalPrice() int
}

type ShoppingCart struct {
	Products []Product
}

func (c *ShoppingCart) AddProduct(item Product) {
	c.Products = append(c.Products, item)
}

func (c *ShoppingCart) GetTotalPrice() int {
	var total int = 0
	for _, item := range c.Products {
		total += item.Price
	}

	return total
}

type Orderer interface {
	Checkout() string
}

type Order struct {
	CustomerName string
	Cart         ShoppingCart
}

func (o *Order) Checkout() string {
	total := o.Cart.GetTotalPrice()
	return fmt.Sprintf("Checkout for %s: %d$", o.CustomerName, total)
}

func main() {
	product1 := Product{Name: "T-Shirt", Price: 25}
	product2 := Product{Name: "Costume", Price: 125}

	//myProducts := []Product{product1, product2}
	//shoppingCart := ShoppingCart{Products: myProducts}

	shoppingCart := ShoppingCart{}

	shoppingCart.AddProduct(product1)
	shoppingCart.AddProduct(product2)

	order := Order{CustomerName: "Alice", Cart: shoppingCart}
	fmt.Println(order.Checkout())

	order2 := Order{CustomerName: "John Doe", Cart: ShoppingCart{
		Products: []Product{
			{Name: "Laptop", Price: 999},
			{Name: "Monitor", Price: 250},
		},
	}}
	fmt.Println(order2.Checkout())

	//shoppingCart.AddProduct(Product{Name: "New Car", Price: 9999})
	//fmt.Println(order2.Checkout())
}
