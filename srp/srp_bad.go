package Book1

import (
	"os"
)

type Book struct {
	Title   string
	Content string
}

func (b *Book) GetTitle() string {
	return b.Title

}
func (b *Book) GetContent() string {
	return b.Content
}

func (b *Book) Save(path string) {
	f, _ := os.Create(path)
	f.WriteString(b.GetContent())
	f.Sync()
	f.Close()
}
