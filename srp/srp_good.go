package Book2

import "os"

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

// Separate content saving logic with book information management
type TextSaver struct {
}

func (t *TextSaver) Save (c Content, path string) {
	f, _ := os.Create(path)
	f.WriteString(c)
	f.Sync()
	f.Close()
}

// Both structs have only one reason to change. 


type Magazine struct {
	Title   string
	Content string
	...
}


func main() {

}
