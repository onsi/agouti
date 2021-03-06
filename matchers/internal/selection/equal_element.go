package selection

import (
	"fmt"
	"github.com/onsi/gomega/format"
)

type EqualElementMatcher struct {
	ExpectedSelection interface{}
}

func (m *EqualElementMatcher) Match(actual interface{}) (success bool, err error) {
	actualSelection, ok := actual.(interface {
		EqualsElement(selection interface{}) (bool, error)
	})

	if !ok {
		return false, fmt.Errorf("EqualElement matcher requires a Selection.  Got:\n%s", format.Object(actual, 1))
	}

	same, err := actualSelection.EqualsElement(m.ExpectedSelection)
	if err != nil {
		return false, fmt.Errorf("EqualElement matcher failed to compare Selections: %s", err)
	}

	return same, nil
}

func (m *EqualElementMatcher) FailureMessage(actual interface{}) (message string) {
	return binarySelectorMessage(actual, "to equal element", m.ExpectedSelection)
}

func (m *EqualElementMatcher) NegatedFailureMessage(actual interface{}) (message string) {
	return binarySelectorMessage(actual, "not to equal element", m.ExpectedSelection)
}
