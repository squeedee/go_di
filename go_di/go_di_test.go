package go_di_test

import (
	. "github.com/squeedee/go_di/go_di"
	"github.com/squeedee/go_di/widget"
	"github.com/squeedee/go_di/widget/fakes"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("GoDi", func() {

	Describe("Unit test", func() {
		var (
			fakeRunner fakes.FakeRunner
			originalNewWidget func(property string) widget.Runner
		)

		BeforeEach(func() {
			fakeRunner.RunReturns("fake run")

			originalNewWidget = widget.NewWidget
			widget.NewWidget = func(property string) widget.Runner {
				return &fakeRunner
			}
		})

		AfterEach(func() {
			widget.NewWidget = originalNewWidget
		})

		It("calls run", func() {
			output := GoDi()

			Expect(output).To(Equal("fake run"))
			Expect(fakeRunner.RunCallCount()).To(Equal(1))
		})
	})

	Describe("Integration test", func() {
		It("has the aceness property", func() {
			Expect(GoDi()).To(Equal("This widget has the property: aceness\n"))
		})
	})

})
