package pkg

type BackstageItem struct{
	item *Item
}

func (b *BackstageItem) Degrade() {
	if b.item.quality < 50 {
		b.item.quality = b.item.quality + 1
		if b.item.sellIn < 11 {
			if b.item.quality < 50 {
				b.item.quality = b.item.quality + 1
			}
		}

		if b.item.sellIn < 6 {
			if b.item.quality < 50 {
				b.item.quality = b.item.quality + 1
			}
		}
	}
	b.item.sellIn = b.item.sellIn - 1

	if b.item.sellIn < 0 {
		b.item.quality = b.item.quality - b.item.quality
	}}
