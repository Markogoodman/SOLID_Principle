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

type WebPage struct {
	Title   string
	Content string
	URL     string
}

// Separate content saving logic with book information management
type TextSaver struct {
}

func (t *TextSaver) Save(content string, path string) {
	f, _ := os.Create(path)
	f.WriteString(content)
	f.Sync()
	f.Close()
}

func main() {
	t := TextSaver{}
	w := WebPage{"title", "content", "url"}
	t.Save(w.Content, "./output.txt")
}
