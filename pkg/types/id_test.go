package types_test

import (
	"code.cestus.io/libs/gotypes/pkg/types"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Id", func() {
	Context("FromString", func() {
		It("should return the id when it is a valid id", func() {
			p := types.DefaultIDProvider
			var id types.GenID
			str := "deadbeef-dead-beef-dead-beef00000075"
			err := p.FromString(&id, str)
			Expect(err).ToNot(HaveOccurred())
			Expect(id.String()).To(Equal(str))
		})
		It("should return an error if the id is not valid", func() {
			p := types.DefaultIDProvider
			var id types.GenID
			str := "deadbeef-dead-beef-dead-beef000000758"
			err := p.FromString(&id, str)
			Expect(err).To(HaveOccurred())
		})
	})

})
