package main

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Metadata struct {
	ID         string
	CreatedAt  time.Time
	ModifiedAt time.Time
}

// New initializes metadata fields
func (m *Metadata) New() {
	m.ID = uuid.New().String()
	m.CreatedAt = time.Now()
	m.ModifiedAt = m.CreatedAt
}

// Customer.New() uses the promoted Metadata.New() method.
// Calling Customer.New() will initialize Customer.Metadata, but
// will not modify Customer specific fields.
type Customer struct {
	Metadata
	Name string
}

// Product.New(string) shadows `Metadata.New() method. You cannot
// call `Product.New()`, but call `Product.New(string)` or
// `Product.Metadata.New()`
type Product struct {
	Metadata
	SKU string
}

func (p *Product) New(sku string) {
	// Initialize the metadata part of product
	p.Metadata.New()
	p.SKU = sku
}

func main() {
	c := Customer{}
	c.New() // Initialize customer metadata
	fmt.Printf("Customer: %+v\n", c)

	p := Product{}
	p.New("sku") // Initialize product metadata and sku
	fmt.Printf("Product: %+v\n", p)

}
