package pkg

type ConjuredItem struct {
	item *Item
}

func (i ConjuredItem) Degrade() {
	if i.item.quality > 0 {
		i.item.quality = i.item.quality - 2
	}
	i.item.sellIn = i.item.sellIn - 1

	if i.item.sellIn < 0 {
		if i.item.quality > 0 {
			i.item.quality = i.item.quality - 2
		}
	}}

