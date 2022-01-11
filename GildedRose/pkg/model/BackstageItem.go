package model

type BackstageItem struct {
	Item *Item
}

func (b *BackstageItem) Degrade() {
	if b.Item.Quality < 50 {
		b.Item.Quality = b.Item.Quality + 1
		if b.Item.SellIn < 11 {
			if b.Item.Quality < 50 {
				b.Item.Quality = b.Item.Quality + 1
			}
		}

		if b.Item.SellIn < 6 {
			if b.Item.Quality < 50 {
				b.Item.Quality = b.Item.Quality + 1
			}
		}
	}
	b.Item.SellIn = b.Item.SellIn - 1

	if b.Item.SellIn < 0 {
		b.Item.Quality = b.Item.Quality - b.Item.Quality
	}
}
