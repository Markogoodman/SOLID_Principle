package Book2

import "os"

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
}

// what if we need a HTMLSaver
// Book struct has to be modified manually
// or add some more code to TextSaver
type Book struct {
	HTMLSaver
	Title   string
	Content string
}
