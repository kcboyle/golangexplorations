package matriarch_test

import (
	"forest/matriarch"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Denizen", func() {
	It("should have an item to trade", func() {
		den := matriarch.NewDenizen()
		Expect(den.TradeItem).ToNot(Equal(""))
	})

	It("should have a random item from list to trade", func() {
		den1 := matriarch.NewDenizen()
		den2 := matriarch.NewDenizen()
		Expect(den1.TradeItem).NotTo(Equal(den2.TradeItem))
	})
})
