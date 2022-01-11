package model

type AgedBrieItem struct {
	Item *Item
}

func (i *AgedBrieItem) Degrade() {
	if i.Item.Quality < 50 {
		i.Item.Quality = i.Item.Quality + 1
	}
	i.Item.SellIn = i.Item.SellIn - 1

	if i.Item.SellIn < 0 {
		if i.Item.Quality < 50 {
			i.Item.Quality = i.Item.Quality + 1
		}
	}
}
