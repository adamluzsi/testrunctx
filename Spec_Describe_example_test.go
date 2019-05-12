package testcase_test

import (
	"github.com/stretchr/testify/require"
	"testing"

	"github.com/adamluzsi/testcase"
)

func ExampleSpec_Describe(t *testing.T) {
	s := testcase.NewSpec(t)

	myType := func(t *testcase.T) *MyType {
		return &MyType{Field1: `input`}
	}

	s.Describe(`IsLower`, func(s *testcase.Spec) {
		subject := func(t *testcase.T) bool { return myType(t).IsLower() }

		s.Then(`test-case`, func(t *testcase.T) {
			require.True(t, subject(t))
		})
	})
}