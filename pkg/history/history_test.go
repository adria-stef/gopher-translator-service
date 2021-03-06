package history_test

import (
	. "github.com/adria-stef/gopher-translator-service/pkg/history"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("History", func() {

	It("adds successfully", func() {
		History().Add("batman", "robin")
		Expect(History()["batman"]).To(Equal("robin"))
	})

	It("arranges properly", func() {
		History().Add("batman", "robin")
		History().Add("antman", "wasp")
		History().Add("Ironman", "Pepper Potts")

		Expect(History().GetArranged()).To(Equal([]map[string]string{
			{"antman": "wasp"},
			{"batman": "robin"},
			{"Ironman": "Pepper Potts"},
		}))
	})

})
