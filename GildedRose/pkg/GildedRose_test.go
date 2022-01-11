package pkg

import (
	"GildedRose/pkg/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_should_Normal_item_decrease_quality_by_one_when_sell_has_not_passed_out(t *testing.T) {
	givenQuality := 10
	givenSellIn := 10
	item := &model.Item{Name: "AnItem", Quality: givenQuality, SellIn: givenSellIn}
	gildedRose := GildedRose{items: []*model.Item{item}}

	gildedRose.UpdateQuality()

	assert.Equal(t, givenQuality-1, item.Quality)
	assert.Equal(t, givenSellIn-1, item.SellIn)
}

func Test_should_Normal_item_decrease_quality_twice_as_fast_when_sell_passed_out(t *testing.T) {
	givenQuality := 10
	item := &model.Item{Name: "AnItem", Quality: givenQuality, SellIn: 0}
	gildedRose := GildedRose{items: []*model.Item{item}}

	gildedRose.UpdateQuality()

	assert.Equal(t, givenQuality-2, item.Quality)
	assert.Equal(t, -1, item.SellIn)
}

func Test_should_Normal_item_not_decrease_quality_when_quality_is_zero(t *testing.T) {
	zeroQuality := 0
	item := &model.Item{Name: "AnItem", Quality: zeroQuality, SellIn: 0}
	gildedRose := GildedRose{items: []*model.Item{item}}

	gildedRose.UpdateQuality()

	assert.Equal(t, zeroQuality, item.Quality)
	assert.Equal(t, -1, item.SellIn)
}

func Test_should_Aged_Brie_increase_quality_when_sell_in_decrease(t *testing.T) {
	givenQuality := 10
	givenSellIn := 10
	item := &model.Item{Name: "Aged Brie", Quality: givenQuality, SellIn: givenSellIn}
	gildedRose := GildedRose{items: []*model.Item{item}}

	gildedRose.UpdateQuality()

	assert.Equal(t, givenQuality+1, item.Quality)
	assert.Equal(t, givenSellIn-1, item.SellIn)
}

func Test_should_not_quality_increase_more_than_50(t *testing.T) {
	givenQuality := 50
	givenSellIn := 10
	item := &model.Item{Name: "Aged Brie", Quality: givenQuality, SellIn: givenSellIn}
	gildedRose := GildedRose{items: []*model.Item{item}}

	gildedRose.UpdateQuality()

	assert.Equal(t, givenQuality, item.Quality)
	assert.Equal(t, givenSellIn-1, item.SellIn)
}

func Test_should_Sulfuras_not_decrease_quality(t *testing.T) {
	givenQuality := 50
	givenSellIn := 10
	item := &model.Item{Name: "Sulfuras, Hand of Ragnaros", Quality: givenQuality, SellIn: givenSellIn}
	gildedRose := GildedRose{items: []*model.Item{item}}

	gildedRose.UpdateQuality()

	assert.Equal(t, givenQuality, item.Quality)
	assert.Equal(t, givenSellIn, item.SellIn)
}

func Test_should_Backstage_increase_quality_by_1_when_SellIn_is_bigger_than_10(t *testing.T) {
	givenQuality := 10
	givenSellIn := 11
	item := &model.Item{Name: "Backstage passes to a TAFKAL80ETC concert", Quality: givenQuality, SellIn: givenSellIn}
	gildedRose := GildedRose{items: []*model.Item{item}}

	gildedRose.UpdateQuality()

	assert.Equal(t, givenQuality+1, item.Quality)
	assert.Equal(t, givenSellIn-1, item.SellIn)
}

func Test_should_Backstage_increase_quality_by_2_when_SellIn_is_below_10(t *testing.T) {
	givenQuality := 10
	givenSellIn := 6
	item := &model.Item{Name: "Backstage passes to a TAFKAL80ETC concert", Quality: givenQuality, SellIn: givenSellIn}
	gildedRose := GildedRose{items: []*model.Item{item}}

	gildedRose.UpdateQuality()

	assert.Equal(t, givenQuality+2, item.Quality)
	assert.Equal(t, givenSellIn-1, item.SellIn)
}

func Test_should_Backstage_increase_quality_by_3_when_SellIn_is_below_5(t *testing.T) {
	givenQuality := 10
	givenSellIn := 5
	item := &model.Item{Name: "Backstage passes to a TAFKAL80ETC concert", Quality: givenQuality, SellIn: givenSellIn}
	gildedRose := GildedRose{items: []*model.Item{item}}

	gildedRose.UpdateQuality()

	assert.Equal(t, givenQuality+3, item.Quality)
	assert.Equal(t, givenSellIn-1, item.SellIn)
}

func Test_should_Backstage_quality_drop_to_0_when_SellIn_is_0(t *testing.T) {
	givenQuality := 10
	givenSellIn := 0
	item := &model.Item{Name: "Backstage passes to a TAFKAL80ETC concert", Quality: givenQuality, SellIn: givenSellIn}
	gildedRose := GildedRose{items: []*model.Item{item}}

	gildedRose.UpdateQuality()

	assert.Equal(t, 0, item.Quality)
	assert.Equal(t, givenSellIn-1, item.SellIn)
}

func Test_should_decrease_twice_his_quality_when_sellIn_decrease(t *testing.T) {
	givenQuality := 10
	givenSellIn := 10
	item := &model.Item{Name: "Conjured", Quality: givenQuality, SellIn: givenSellIn}
	gildedRose := GildedRose{items: []*model.Item{item}}

	gildedRose.UpdateQuality()

	assert.Equal(t, givenQuality-2, item.Quality)
	assert.Equal(t, givenSellIn-1, item.SellIn)
}

func Test_should_decrease_four_times_his_quality_when_sellIn_is_0(t *testing.T) {
	givenQuality := 10
	givenSellIn := 0
	item := &model.Item{Name: "Conjured", Quality: givenQuality, SellIn: givenSellIn}
	gildedRose := GildedRose{items: []*model.Item{item}}

	gildedRose.UpdateQuality()

	assert.Equal(t, givenQuality-4, item.Quality)
	assert.Equal(t, givenSellIn-1, item.SellIn)
}
