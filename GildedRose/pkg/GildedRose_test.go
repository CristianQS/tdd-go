package pkg

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_should_decrease_quality_twice_as_fast_when_sell_passed_out(t *testing.T) {
	givenQuality := 10
	item := &Item{name: "a", quality: givenQuality, sellIn: 0}
	gildedRose := GildedRose{items: []*Item{item}}

	gildedRose.UpdateQuality()

	assert.Equal(t, givenQuality,item.quality)
}

