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

type TextSaver struct {
}

func (t *TextSaver) Save(c Content, path string) {
	f, _ := os.Create(path)
	f.WriteString(c)
	f.Sync()
	f.Close()
}

type Magazine struct {
	Title   string
	Content string
	...
}


func main() {

}
