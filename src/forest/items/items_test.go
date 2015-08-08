package items_test

import (
	"forest/items"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Items", func() {
	It("should populate the allItems structure", func() {
		item := items.NewItems()
		Expect(item.GetAllItems()).ToNot(BeEmpty())
		Expect(item.GetAllItems()).To(ContainElement("flower"))
	})

	It("should populate the tradeable items channel", func() {
		its := items.NewItems()
	})
})
