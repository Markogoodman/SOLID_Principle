package main

type Item interface {
	GetPrice()
}

type Book struct {
	Title string
	Pages int
}

func (b *Book) GetPrice() int {
	return b.Pages * 5
}

type Magazine struct {
	Title string
	Pages int
}

func (m *Magazine) GetPrice() int {
	return m.Pages * 3
}

type PriceCalculator struct {
}

// return price sum of all items
func (pc *PriceCalculator) Sum(items []Item) int {
	sum := 0
	// New item can be added without modifying this method
	for _, item := range items {
		sum += item.GetPrice()
	}
	return sum
}
