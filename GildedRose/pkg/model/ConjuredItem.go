package model

type ConjuredItem struct {
	Item *Item
}

func (i ConjuredItem) Degrade() {
	if i.Item.Quality > 0 {
		i.Item.Quality = i.Item.Quality - 2
	}
	i.Item.SellIn = i.Item.SellIn - 1

	if i.Item.SellIn < 0 {
		if i.Item.Quality > 0 {
			i.Item.Quality = i.Item.Quality - 2
		}
	}
}
