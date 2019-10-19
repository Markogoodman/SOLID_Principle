package Bad

type Book struct {
	Title string
	Pages int
}

func (b *Book) GetPages() int {
	return b.Pages
}

type Magazine struct {
	Title string
	Pages int
}

func (m *Magazine) GetPages() int {
	return m.Pages
}

type PriceCalculator struct {
}

// return price sum of all items
func (pc *PriceCalculator) Sum(items []interface{}) int {
	sum := 0
	for _, item := range items {
		switch i := item.(type) {
		case Book:
			sum += i.GetPages() * 5
		case Magazine:
			sum += i.GetPages() * 3
		}
		// Bad: once you have new item, this function should be modified
	}
	return sum
}
