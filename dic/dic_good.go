package Book2

import "os"

type ContentSaver interface {
	Save(content string, path string)
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

type HTMLSaver struct {
}

func (h *HTMLSaver) Save(content string, path string) {
	f, _ := os.Create(path)
	f.WriteString("<html>" + content + "</html>")
	f.Sync()
	f.Close()
}

type Book struct {
	Saver   ContentSaver
	Title   string
	Content string
}

// Pass saver as a parameter when initialization.
func NewBook(saver ContentSaver) *Book {
	book := new(Book)
	book.Saver = saver
	return book
}
