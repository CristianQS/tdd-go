package pkg

type GildedRose struct {
	items []*Item
}

func (g *GildedRose) UpdateQuality() {
	for i := 0; i < len(g.items); i++ {
		item := create(g.items[i])
		item.Degrade()
	}
}
