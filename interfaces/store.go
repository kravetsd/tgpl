package interfaces

import "fmt"

type Person struct {
	Name    string
	Surname string
}

type Book struct {
	Title  string
	Price  int
	Author Person
	Review string
}

func (book Book) Present() {
	fmt.Printf("Title: %v\tPrice: %v\tDescription: %v\tAuthor: %v\n", book.Title, book.Price, book.Review, book.Author.Name+" "+book.Author.Surname)
}

type Game struct {
	Name        string
	Price       int
	Description string
}

func (game Game) Present() {
	fmt.Printf("Title: %v\tPrice: %v\tDescription: %v\n", game.Name, game.Price, game.Description)
}

type Puzle struct {
	Name        string
	Price       int
	Description string
}

func (puzle Puzle) Present() {
	fmt.Printf("Title: %v\tPrice: %v\tDescription: %v\n", puzle.Name, puzle.Price, puzle.Description)
}

type Seller interface {
	Present()
}

var show []Seller

func Shopping() {
	azbk := Book{Title: "Azbuka", Price: 12, Author: Person{Name: "Ivan", Surname: "Petrov"}, Review: "School book"}
	game := Game{Name: "Tetris", Price: 40, Description: "Interesting computer game."}
	puzle := Puzle{Name: "kube", Price: 30, Description: "Intereseting and popular puzle"}

	show = append(show, azbk, game, puzle)

	for _, v := range show {
		v.Present()
	}
}
