package pkg

import "GildedRose/pkg/model"

func create(item *model.Item) model.DegradableItem {
	switch item.Name {
	case "Sulfuras, Hand of Ragnaros":
		return &model.Sulfuras{Item: item}
	case "Aged Brie":
		return &model.AgedBrieItem{Item: item}
	case "Backstage passes to a TAFKAL80ETC concert":
		return &model.BackstageItem{Item: item}
	case "Conjured":
		return &model.ConjuredItem{Item: item}
	default:
		return &model.NormalItem{Item: item}
	}
}
