package rag_test

import (
	"testing"

	"github.com/thehungry-dev/rag/pkg/ctrls"
)

func TestTagFilterSelect(t *testing.T) {
	t.Parallel()
	tagFilter := ctrls.TagFilterExample()

	t.Run("Selects no excluded tags", func(t *testing.T) {
		tags := ctrls.TagsMatchingExample()

		selected := tagFilter.Select(tags)

		if !selected {
			t.Error()
		}

	})

	t.Run("Rejects excluded tags", func(t *testing.T) {
		tags := ctrls.TagsExcludedMatchingExample()

		rejected := tagFilter.Reject(tags)

		if !rejected {
			t.Error()
		}
	})

	t.Run("Selects required and excluded tags", func(t *testing.T) {
		tagFilter = ctrls.TagFilterRequiredExample()
		tags := ctrls.TagsMatchingExample()

		selected := tagFilter.Select(tags)

		if !selected {
			t.Error()
		}
	})
}
