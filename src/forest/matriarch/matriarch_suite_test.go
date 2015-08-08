package matriarch_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestMatriarch(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Matriarch Suite")
}
