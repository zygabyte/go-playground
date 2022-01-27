package main

import (
	"fmt"
)

type Order interface {
	getName() string
	getDescription() string
	getPrice(withDiscount bool) float32
}

// Product implementing order with implicit interfaces
type Product struct {
	name, description string
	price, discount   float32
}

func (p Product) getName() string {
	return p.name
}

func (p Product) getDescription() string {
	return p.description
}

func (p Product) getPrice(withDiscount bool) float32 {
	if withDiscount {
		return p.price * p.discount
	}
	return p.price
}

type ProductList []Product // ProductList is called a type alias, representing a slice of Products

func (p ProductList) ViewProducts() {
	for _, product := range p {
		fmt.Printf("%v", product.name)
	}
}

// Service implementing order with implicit interfaces
type Service struct {
	name, information string
	amount            float32
}

func (p Service) getName() string {
	return p.name
}

func (p Service) getDescription() string {
	return p.information
}

func (p Service) getPrice(_ bool) float32 {
	return p.amount
}

func main() {
	orders := []Order{
		Product{"t-shirt", "black and white", 300, 1.7},
		Service{"airtime", "subscribe monthly", 500},
	}

	for _, v := range orders {
		fmt.Printf("value %v %v %v \n", v.getName(), v.getDescription(), v.getPrice(true))
	}

	fmt.Printf("\n\ntesting type alias of products \n\n")
	products := ProductList{
		Product{"t-shirt", "black and white", 300, 1.7},
		Product{"i-phone", "cool phone", 20000, 0.5},
	}

	for _, product := range products {
		fmt.Printf("product -> %v \ndescription -> %v \nprice -> %v \ndiscount -> %v \n\n", product.name, product.description, product.price, product.discount)
	}
}
