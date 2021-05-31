package pkg

import "fmt"

type Item struct {
	name    string
	sellIn  int
	quality int
}

func (i *Item) ToString() string {
	return fmt.Sprintf("%s, %v, %v", i.name, i.sellIn, i.quality)
}
