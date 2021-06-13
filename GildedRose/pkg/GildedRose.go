package pkg

type GildedRose struct {
	items []*Item
}

func (g *GildedRose) UpdateQuality() {
	for i := 0; i < len(g.items); i++ {
		item := create(g.items[i])
		item.Degrade()

		//if g.items[i].name == "Aged Brie" {
		//	AgedBrieUpdate(g.items[i])
		//} else if g.items[i].name == "Backstage passes to a TAFKAL80ETC concert" {
		//	BackstageUpdate(g.items[i])
		//} else {
		//	NormalItemUpdate(g.items[i])
		//}
	}
}

func NormalItemUpdate(item *Item) {
	if item.quality > 0 {
		item.quality = item.quality - 1
	}
	item.sellIn = item.sellIn - 1

	if item.sellIn < 0 {
		if item.quality > 0 {
			item.quality = item.quality - 1
		}
	}
}

func BackstageUpdate(item *Item) {
	if item.quality < 50 {
		item.quality = item.quality + 1
		if item.sellIn < 11 {
			if item.quality < 50 {
				item.quality = item.quality + 1
			}
		}

		if item.sellIn < 6 {
			if item.quality < 50 {
				item.quality = item.quality + 1
			}
		}
	}
	item.sellIn = item.sellIn - 1

	if item.sellIn < 0 {
		item.quality = item.quality - item.quality
	}
}

func AgedBrieUpdate(item *Item) {
	if item.quality < 50 {
		item.quality = item.quality + 1
	}
	item.sellIn = item.sellIn - 1

	if item.sellIn < 0 {
		if item.quality < 50 {
			item.quality = item.quality + 1
		}
	}
}
