package Bad

type Book struct {
	Title string
	Pages int
}

type Magazine struct {
	Title string
	Pages int
}

type PriceCalculator struct {
}

// return price sum of all items
func (pc *PriceCalculator) Sum(items []interface{}) int {
	sum := 0
	for _, item := range items {
		switch i := item.(type) {
		case Book:
			sum += i.Pages * 5
		case Magazine:
			sum += i.Pages * 3
		}
		// Bad: once you have a new item, this function should be modified
	}
	return sum
}
