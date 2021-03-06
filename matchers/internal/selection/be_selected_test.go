package selection_test

import (
	"github.com/sclevine/agouti/matchers/internal/mocks"
	. "github.com/sclevine/agouti/matchers/internal/selection"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("BeSelectedMatcher", func() {
	var (
		matcher   *BeSelectedMatcher
		selection *mocks.Selection
	)

	BeforeEach(func() {
		selection = &mocks.Selection{}
		selection.StringCall.ReturnString = "CSS: #selector"
		matcher = &BeSelectedMatcher{}
	})

	Describe("#Match", func() {
		Context("when the actual object is a selection", func() {
			Context("when the element is selected", func() {
				BeforeEach(func() {
					selection.SelectedCall.ReturnSelected = true
				})

				It("returns true", func() {
					success, _ := matcher.Match(selection)
					Expect(success).To(BeTrue())
				})

				It("does not return an error", func() {
					_, err := matcher.Match(selection)
					Expect(err).NotTo(HaveOccurred())
				})
			})

			Context("when the element is not selected", func() {
				BeforeEach(func() {
					selection.SelectedCall.ReturnSelected = false
				})

				It("returns false", func() {
					success, _ := matcher.Match(selection)
					Expect(success).To(BeFalse())
				})

				It("does not return an error", func() {
					_, err := matcher.Match(selection)
					Expect(err).NotTo(HaveOccurred())
				})
			})
		})

		Context("when the actual object is not a selection", func() {
			It("returns an error", func() {
				_, err := matcher.Match("not a selection")
				Expect(err).To(MatchError("BeSelected matcher requires a Selection.  Got:\n    <string>: not a selection"))
			})
		})
	})

	Describe("#FailureMessage", func() {
		It("returns a failure message", func() {
			selection.SelectedCall.ReturnSelected = false
			matcher.Match(selection)
			message := matcher.FailureMessage(selection)
			Expect(message).To(Equal("Expected selection 'CSS: #selector' to be selected"))
		})
	})

	Describe("#NegatedFailureMessage", func() {
		It("returns a negated failure message", func() {
			selection.SelectedCall.ReturnSelected = true
			matcher.Match(selection)
			message := matcher.NegatedFailureMessage(selection)
			Expect(message).To(Equal("Expected selection 'CSS: #selector' not to be selected"))
		})
	})
})
