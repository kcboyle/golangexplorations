package items_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestItems(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Items Suite")
}
