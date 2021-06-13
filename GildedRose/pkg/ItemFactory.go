package pkg

func create(item *Item) DegradableItem{
	switch item.name {
	case "Sulfuras, Hand of Ragnaros":
		return &Sulfuras{item: item}
	case "Aged Brie":
		return &AgedBrieItem{item: item}
	case "Backstage passes to a TAFKAL80ETC concert":
		return &BackstageItem{item: item}
	default:
		return &NormalItem{item: item}
	}
}

