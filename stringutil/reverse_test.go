package stringutil_test

import (
	"ginkgoreverse/stringutil"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ginkgoreverse", func() {

	var value string
	Context("Here we reverse the string", func() {

		BeforeEach(func() {
			value = stringutil.Reverse("hello")
		})

		It("Successfully reveresed", func() {
			Expect(value).To(Equal("olleh"))
		})

	})

})
