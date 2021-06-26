package magic_test

import (
	"io/ioutil"

	"github.com/ralch/magic"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("DetectContentType", func() {
	Context("when the document is csv file", func() {
		It("returns the mime type", func() {
			data, err := ioutil.ReadFile("./fixture/document.csv")
			Expect(err).NotTo(HaveOccurred())
			Expect(magic.DetectContentType(data)).To(Equal("text/csv"))
		})
	})

	Context("when the document is excel file", func() {
		It("returns the mime type", func() {
			data, err := ioutil.ReadFile("./fixture/document.xlsx")
			Expect(err).NotTo(HaveOccurred())
			Expect(magic.DetectContentType(data)).To(Equal("application/vnd.openxmlformats-officedocument.spreadsheetml.sheet"))
		})
	})
})
