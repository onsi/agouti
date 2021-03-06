package selection_test

import (
	"github.com/sclevine/agouti/matchers/internal/mocks"
	. "github.com/sclevine/agouti/matchers/internal/selection"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("HaveAttributeMatcher", func() {
	var (
		matcher   *HaveAttributeMatcher
		selection *mocks.Selection
	)

	BeforeEach(func() {
		selection = &mocks.Selection{}
		selection.StringCall.ReturnString = "CSS: #selector"
		matcher = &HaveAttributeMatcher{ExpectedAttribute: "some-attribute", ExpectedValue: "some value"}
	})

	Describe("#Match", func() {
		Context("when the actual object is a selection", func() {
			It("requests the provided page attribute", func() {
				matcher.Match(selection)
				Expect(selection.AttributeCall.Attribute).To(Equal("some-attribute"))
			})

			Context("when the expected attribute value matches the actual attribute value", func() {
				BeforeEach(func() {
					selection.AttributeCall.ReturnValue = "some value"
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

			Context("when the expected attribute value does not match the actual attribute value", func() {
				BeforeEach(func() {
					selection.AttributeCall.ReturnValue = "some other value"
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
				Expect(err).To(MatchError("HaveAttribute matcher requires a Selection.  Got:\n    <string>: not a selection"))
			})
		})
	})

	Describe("#FailureMessage", func() {
		It("returns a failure message", func() {
			selection.AttributeCall.ReturnValue = "some other value"
			matcher.Match(selection)
			message := matcher.FailureMessage(selection)
			Expect(message).To(ContainSubstring("Expected selection 'CSS: #selector' to have attribute matching\n    [some-attribute=\"some value\"]"))
			Expect(message).To(ContainSubstring("but found\n    [some-attribute=\"some other value\"]"))
		})
	})

	Describe("#NegatedFailureMessage", func() {
		It("returns a negated failure message", func() {
			selection.AttributeCall.ReturnValue = "some value"
			matcher.Match(selection)
			message := matcher.NegatedFailureMessage(selection)
			Expect(message).To(ContainSubstring("Expected selection 'CSS: #selector' not to have attribute matching\n    [some-attribute=\"some value\"]"))
			Expect(message).To(ContainSubstring("but found\n    [some-attribute=\"some value\"]"))
		})
	})
})
