package book

import "fmt"

type Printable interface {
	PrintInfo()
}

func Print(p Printable) {
	p.PrintInfo()
}

func (t TextBook) PrintInfo() {
	fmt.Printf("Title: %s\nAuthor: %s\nPages: %d\nEditorial:%s\nNivel:%s\n\n", t.title, t.author, t.pages, t.editorial, t.level)
}

type Book struct {
	Title  string
	Author string
	Pages  int
} //Modelo del objeto de el libro

type PrivateBook struct {
	title  string
	author string
	pages  int
} //Modelo del objeto de el libro privado

func (b Book) PrintInfo() {
	fmt.Printf("Title: %s\nAuthor: %s\nPages: %d\n",
		b.Title, b.Author, b.Pages)
}

func (b *PrivateBook) PrintPrivateInfo() {
	fmt.Printf("Title: %s\nAuthor: %s\nPages: %d\n",
		b.title, b.author, b.pages)
}

func (b *PrivateBook) SetTitle(title string) {
	b.title = title
}

func (b *PrivateBook) GetTitle() string {
	return b.title
}

func NewBook(title string, author string, pages int) *Book {
	return &Book{
		Title:  title,
		Author: author,
		Pages:  pages}
}

func NewPrivateBook(title string, author string, pages int) *PrivateBook {
	return &PrivateBook{
		title:  title,
		author: author,
		pages:  pages}
}

type TextBook struct {
	PrivateBook
	editorial string
	level     string
}

func NewTextBook(title, author string, pages int, editorial, level string) *TextBook {
	return &TextBook{
		PrivateBook: PrivateBook{title, author, pages},
		editorial:   editorial,
		level:       level}
}
