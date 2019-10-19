package main

import (
	"os"
)

type Book struct {
	Title     string
	Content   string
	Author    string
	Publisher string
	Price     float32
}

// Manage book information
func (b *Book) GetTitle() string {
	return b.Title

}
func (b *Book) GetContent() string {
	return b.Content
}

// Save book content to disk
func (b Book) Save(content string, path string) {
	f, _ := os.Create(path)
	f.WriteString(content)
	f.Sync()
	f.Close()
}

// What if WebPage needs to save content.
// WebPage depends on some unnecessary fields (Book.Price Book.Author ...).
type WebPage struct {
	Book
	Title   string
	Content string
	URL     string
}

func main() {
	// WebPage depends on some unnecessary fields (Book.Price Book.Author ...).

	m := WebPage{Book{}, "title", "content", "url"}
	m.Save(m.Content, "./output.txt")

	// m := WebPage{"title", "content", "url"}
	// b := Book{...}
	// b.Save(m.Content, "./output.txt")

}
