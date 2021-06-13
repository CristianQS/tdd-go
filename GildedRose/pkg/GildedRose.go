package pkg

type GildedRose struct {
	items []*Item
}

func (g *GildedRose) UpdateQuality() {
	for i := 0; i < len(g.items); i++ {
		if g.items[i].name == "Aged Brie" {
			AgedBrieUpdate(g.items[i])
		}
		if g.items[i].name == "Backstage passes to a TAFKAL80ETC concert" {
			BackstageUpdate(g.items[i])
		}
		if !(g.items[i].name == "Aged Brie") && !(g.items[i].name == "Backstage passes to a TAFKAL80ETC concert") {
			if g.items[i].quality > 0 {
				if !(g.items[i].name == "Sulfuras, Hand of Ragnaros") {
					g.items[i].quality = g.items[i].quality - 1
				}
			}
		}

		if !(g.items[i].name == "Sulfuras, Hand of Ragnaros") && !(g.items[i].name == "Aged Brie") &&
			!(g.items[i].name == "Backstage passes to a TAFKAL80ETC concert") {
			g.items[i].sellIn = g.items[i].sellIn - 1
		}

		if g.items[i].sellIn < 0 {
			if !(g.items[i].name == "Aged Brie") {
				if !(g.items[i].name == "Backstage passes to a TAFKAL80ETC concert") {
					if g.items[i].quality > 0 {
						if !(g.items[i].name == "Sulfuras, Hand of Ragnaros") {
							g.items[i].quality = g.items[i].quality - 1
						}
					}
				} else {
				}
			} else {

			}
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
