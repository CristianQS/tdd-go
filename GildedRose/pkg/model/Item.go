package model

import "fmt"

type Item struct {
	Name    string
	SellIn  int
	Quality int
}

func (i *Item) ToString() string {
	return fmt.Sprintf("%s, %v, %v", i.Name, i.SellIn, i.Quality)
}
