package pkg

type GildedRose struct {
	items []*Item
}

func (g *GildedRose) UpdateQuality() {
	for i := 0; i < len(g.items); i++ {
		if !IsAgedBrie(g, i) && !IsBackstage(g, i) {
			if !(g.items[i].name == "Sulfuras, Hand of Ragnaros") {
				NormalItemUpdate(g.items[i])
			}
		} else {
			if g.items[i].quality < 50 {
				AgedBrieItemUpdate(g.items[i])

				if IsBackstage(g, i) {
					BackstageItemUpdate(g.items[i])
				}
			}
		}

		if !(g.items[i].name == "Sulfuras, Hand of Ragnaros") {
			g.items[i].sellIn = g.items[i].sellIn - 1
		}

		if g.items[i].sellIn < 0 {
			if !IsAgedBrie(g, i) {
				if !IsBackstage(g, i) {
					if g.items[i].quality > 0 {
						if !(g.items[i].name == "Sulfuras, Hand of Ragnaros") {
							g.items[i].quality = g.items[i].quality - 1
						}
					}
				} else {
					g.items[i].quality = g.items[i].quality - g.items[i].quality
				}
			} else {
				if g.items[i].quality < 50 {
					g.items[i].quality = g.items[i].quality + 1
				}
			}
		}
	}
}

func IsBackstage(g *GildedRose, i int) bool {
	return g.items[i].name == "Backstage passes to a TAFKAL80ETC concert"
}

func IsAgedBrie(g *GildedRose, i int) bool {
	return g.items[i].name == "Aged Brie"
}

func NormalItemUpdate(item *Item) {
	if item.quality > 0 {
		item.quality = item.quality - 1
	}
}

func AgedBrieItemUpdate(item *Item) {
	item.quality = item.quality + 1
}

func BackstageItemUpdate(item *Item) {
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
