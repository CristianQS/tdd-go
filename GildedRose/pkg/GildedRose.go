package pkg

import "GildedRose/pkg/model"

type GildedRose struct {
	items []*model.Item
}

func (g *GildedRose) UpdateQuality() {
	for _, item := range g.items {
		degradableItem := create(item)
		degradableItem.Degrade()
	}
}
