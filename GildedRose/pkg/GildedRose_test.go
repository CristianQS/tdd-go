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
}

func Test_should_not_decrease_quality_when_quality_is_zero(t *testing.T) {
	zeroQuality := 0
	item := &Item{name: "a", quality: zeroQuality, sellIn: 0}
	gildedRose := GildedRose{items: []*Item{item}}

	gildedRose.UpdateQuality()

	assert.Equal(t, zeroQuality, item.quality)
}
func Test_should_Aged_Brie_increase_quality_when_sell_in_decrease(t *testing.T) {
	givenQuality := 10
	item := &Item{name: "a", quality: givenQuality, sellIn: 10}
	gildedRose := GildedRose{items: []*Item{item}}

	gildedRose.UpdateQuality()

	assert.Equal(t, givenQuality, item.quality+1)
}
