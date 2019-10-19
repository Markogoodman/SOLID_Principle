package Book1

import (
	"os"
)

type Book struct {
	Title   string
	Content string
}

// Manage book information
func (b *Book) GetTitle() string {
	return b.Title

}
func (b *Book) GetContent() string {
	return b.Content
}

// Save book content to disk
func (b *Book) Save(path string) {
	f, _ := os.Create(path)
	f.WriteString(b.Content)
	f.Sync()
	f.Close()
}

//  Book struct has to be changed if one of these two logic needs to be changed
