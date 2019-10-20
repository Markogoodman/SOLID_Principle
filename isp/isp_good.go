package main

import (
	"os"
)

// Separate content saving logic with book information management
type TextSaver struct {
}

func (t *TextSaver) Save(content string, path string) {
	f, _ := os.Create(path)
	f.WriteString(content)
	f.Sync()
	f.Close()
}

type Book struct {
	TextSaver
	Title   string
	Content string
	Author  string
	Price   float32
}

// Manage book information
func (b *Book) GetTitle() string {
	return b.Title

}
func (b *Book) GetContent() string {
	return b.Content
}

type WebPage struct {
	TextSaver
	Title   string
	Content string
	URL     string
}

func main() {
	w := WebPage{TextSaver{}, "title", "content", "url"}
	w.Save(w.Content, "./output.txt")
}
