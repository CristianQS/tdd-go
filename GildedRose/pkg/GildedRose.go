package pkg

type GildedRose struct {
	items []*Item
}

func (g *GildedRose) UpdateQuality() {
	for i := 0; i < len(g.items); i++ {
		if !IsAgedBrie(g, i) && !IsBackstage(g, i) {
			if !(IsRagnaros(g.items[i])) {
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

		if !(IsRagnaros(g.items[i])) {
			g.items[i].sellIn = g.items[i].sellIn - 1
		}

		if g.items[i].sellIn < 0 {
			if !IsAgedBrie(g, i) {
				if !IsBackstage(g, i) {
					NormalItemUpdate(g.items[i])
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
func IsRagnaros(item *Item) bool {
	return item.name == "Sulfuras, Hand of Ragnaros"
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
