package pkg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_should_decrease_quality_twice_as_fast_when_sell_passed_out(t *testing.T) {
	givenQuality := 10
	item := &Item{name: "a", quality: givenQuality, sellIn: 0}
	gildedRose := GildedRose{items: []*Item{item}}

	gildedRose.UpdateQuality()

	assert.Equal(t, givenQuality-2, item.quality)
	assert.Equal(t, -1, item.sellIn)
}

func Test_should_not_decrease_quality_when_quality_is_zero(t *testing.T) {
	zeroQuality := 0
	item := &Item{name: "a", quality: zeroQuality, sellIn: 0}
	gildedRose := GildedRose{items: []*Item{item}}

	gildedRose.UpdateQuality()

	assert.Equal(t, zeroQuality, item.quality)
	assert.Equal(t, -1, item.sellIn)

}

func Test_should_Aged_Brie_increase_quality_when_sell_in_decrease(t *testing.T) {
	givenQuality := 10
	givenSellIn := 10
	item := &Item{name: "Aged Brie", quality: givenQuality, sellIn: givenSellIn}
	gildedRose := GildedRose{items: []*Item{item}}

	gildedRose.UpdateQuality()

	assert.Equal(t, givenQuality+1, item.quality)
	assert.Equal(t, givenSellIn-1, item.sellIn)
}

func Test_should_not_quality_increase_more_than_50(t *testing.T) {
	givenQuality := 50
	givenSellIn := 10
	item := &Item{name: "Aged Brie", quality: givenQuality, sellIn: givenSellIn}
	gildedRose := GildedRose{items: []*Item{item}}

	gildedRose.UpdateQuality()

	assert.Equal(t, givenQuality, item.quality)
	assert.Equal(t, givenSellIn-1, item.sellIn)
}

func Test_should_Sulfuras_not_decrease_quality(t *testing.T) {
	givenQuality := 50
	givenSellIn := 10
	item := &Item{name: "Sulfuras, Hand of Ragnaros", quality: givenQuality, sellIn: givenSellIn}
	gildedRose := GildedRose{items: []*Item{item}}

	gildedRose.UpdateQuality()

	assert.Equal(t, givenQuality, item.quality)
	assert.Equal(t, givenSellIn, item.sellIn)
}

func Test_should_Backstage_increase_quality_by_2_when_SellIn_is_between_10_and_6(t *testing.T) {
	givenQuality := 10
	givenSellIn := 10
	item := &Item{name: "Backstage passes to a TAFKAL80ETC concert", quality: givenQuality, sellIn: givenSellIn}
	gildedRose := GildedRose{items: []*Item{item}}

	gildedRose.UpdateQuality()

	assert.Equal(t, givenQuality+2, item.quality)
	assert.Equal(t, givenSellIn-1, item.sellIn)
}
func Test_should_Backstage_increase_quality_by_3_when_SellIn_is_between_5_and_1(t *testing.T) {
	givenQuality := 10
	givenSellIn := 5
	item := &Item{name: "Backstage passes to a TAFKAL80ETC concert", quality: givenQuality, sellIn: givenSellIn}
	gildedRose := GildedRose{items: []*Item{item}}

	gildedRose.UpdateQuality()

	assert.Equal(t, givenQuality+3, item.quality)
	assert.Equal(t, givenSellIn-1, item.sellIn)
}

func Test_should_Backstage_quality_drop_to_0_when_SellIn_is_0(t *testing.T) {
	givenQuality := 10
	givenSellIn := 0
	item := &Item{name: "Backstage passes to a TAFKAL80ETC concert", quality: givenQuality, sellIn: givenSellIn}
	gildedRose := GildedRose{items: []*Item{item}}

	gildedRose.UpdateQuality()

	assert.Equal(t, 0, item.quality)
	assert.Equal(t, givenSellIn-1, item.sellIn)
}
