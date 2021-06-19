package pkg

type GildedRose struct {
	items []*Item
}

func (g *GildedRose) UpdateQuality() {
	for _, item := range g.items {
		degradableItem := create(item)
		degradableItem.Degrade()
	}
}
