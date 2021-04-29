package pkg

type GildedRose struct {
	items []*Item
}

func (g *GildedRose) UpdateQuality()  {
	for i := 0; i < len(g.items); i++ {
		if !(g.items[i].name == "Aged Brie") && !(g.items[i].name == "Backstage passes to a TAFKAL80ETC concert") {
			if g.items[i].quality > 0 {
				if !(g.items[i].name == "Sulfuras, Hand of Ragnaros") {
					g.items[i].quality = g.items[i].quality - 1;
				}
			}
		} else {
			if g.items[i].quality < 50 {
				g.items[i].quality = g.items[i].quality + 1;

				if g.items[i].name == "Backstage passes to a TAFKAL80ETC concert" {
					if g.items[i].sellIn < 11 {
						if g.items[i].quality < 50 {
							g.items[i].quality = g.items[i].quality + 1;
						}
					}

					if g.items[i].sellIn < 6 {
						if g.items[i].quality < 50 {
							g.items[i].quality = g.items[i].quality + 1;
						}
					}
				}
			}
		}

		if !(g.items[i].name == "Sulfuras, Hand of Ragnaros") {
			g.items[i].sellIn = g.items[i].sellIn - 1;
		}

		if g.items[i].sellIn < 0 {
			if !(g.items[i].name == "Aged Brie") {
				if !(g.items[i].name == "Backstage passes to a TAFKAL80ETC concert") {
					if g.items[i].quality > 0 {
						if !(g.items[i].name == "Sulfuras, Hand of Ragnaros") {
							g.items[i].quality = g.items[i].quality - 1;
						}
					}
				} else {
					g.items[i].quality = g.items[i].quality - g.items[i].quality;
				}
			} else {
				if g.items[i].quality < 50 {
					g.items[i].quality = g.items[i].quality + 1;
				}
			}
		}
	}
}
