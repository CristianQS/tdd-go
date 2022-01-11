package model

type NormalItem struct {
	Item *Item
}

func (i *NormalItem) Degrade() {
	if i.Item.Quality > 0 {
		i.Item.Quality = i.Item.Quality - 1
	}
	i.Item.SellIn = i.Item.SellIn - 1

	if i.Item.SellIn < 0 {
		if i.Item.Quality > 0 {
			i.Item.Quality = i.Item.Quality - 1
		}
	}
}
