package pkg

type AgedBrieItem struct{
	item *Item
}

func (i *AgedBrieItem) Degrade() {
	if i.item.quality < 50 {
		i.item.quality = i.item.quality + 1
	}
	i.item.sellIn = i.item.sellIn - 1

	if i.item.sellIn < 0 {
		if i.item.quality < 50 {
			i.item.quality = i.item.quality + 1
		}
	}
}

