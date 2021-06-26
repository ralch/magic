package magic_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestMagic(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Magic Suite")
}
