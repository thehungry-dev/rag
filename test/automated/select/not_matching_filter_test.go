package rag_test

import (
	"testing"

	"github.com/thehungry-dev/rag/pkg/ctrls"
)

func TestTagFilterRejectNotMatching(t *testing.T) {
	t.Parallel()
	tagFilter := ctrls.TagFilterExample()

	t.Run("Rejects missing required tags", func(t *testing.T) {
		tags := ctrls.TagsNonMatchingMissingRequiredExample()

		rejected := tagFilter.Reject(tags)

		if !rejected {
			t.Error()
		}
	})

	t.Run("Rejects missing required one of tags", func(t *testing.T) {
		tags := ctrls.TagsNonMatchingMissingAllRequiredOneOfExample()

		rejected := tagFilter.Reject(tags)

		if !rejected {
			t.Error()
		}
	})
}
