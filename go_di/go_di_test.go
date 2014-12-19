package go_di_test

import (
	. "github.com/squeedee/go_di/go_di"
	"github.com/squeedee/go_di/go_di/fakes"
	"github.com/squeedee/go_di/widget"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("GoDi", func() {

	Describe("Integration test", func() {
		It("has the aceness property", func() {
			Expect(GoDi()).To(Equal("This widget has the property: aceness\n"))
		})
	})

	Describe("Unit test", func() {
		var fakeWidget fakes.FakeWidget

		BeforeEach(func() {
			widget.NewWidget = func(property string) widget.Runner {
				return &fakeWidget
			}
		})

		It("calls run", func() {
			output := GoDi()

			Expect(output).To(Equal("fake run"))
			Expect(fakeWidget.GetWhatHappened()).To(Equal("Ran Run"))
		})
	})

})
