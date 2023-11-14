package rflag_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/meln5674/rflag"
)

var _ = Describe("ParseTag", func() {
	It("should parse an empty string", func() {
		info, err := rflag.ParseTag("")
		Expect(err).ToNot(HaveOccurred())
		Expect(info).To(Equal(rflag.TagInfo{}))
	})
	It("should parse all fields", func() {
		info, err := rflag.ParseTag("name=a,shorthand=b,usage=c,prefix=d")
		Expect(err).ToNot(HaveOccurred())
		Expect(info).To(Equal(rflag.TagInfo{
			Name:      "a",
			Shorthand: "b",
			Usage:     "c",
			Prefix:    "d",
		}))
	})

	It("should handle escapes", func() {
		info, err := rflag.ParseTag("name=abc,,def,,ghi,shorthand=,,jkl,,,,m,,,usage=,,nop,prefix=,,")
		Expect(err).ToNot(HaveOccurred())
		Expect(info).To(Equal(rflag.TagInfo{
			Name:      "abc,def,ghi",
			Shorthand: ",jkl,,m,",
			Usage:     ",nop",
			Prefix:    ",",
		}))
	})
})
