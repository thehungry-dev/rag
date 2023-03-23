package filter_test

import (
	"testing"

	ctrls "github.com/thehungry-dev/rag/ctrls/tag/filter"
)

func TestTagFilterNotMatchingSelect(t *testing.T) {
	t.Run("Tag", func(t *testing.T) {
		t.Run("Filter", func(t *testing.T) {
			t.Run("Select", func(t *testing.T) {
				tagFilter := ctrls.TagFilterExample()

				t.Run("Not matching filter", func(t *testing.T) {
					t.Run("Missing required tags", func(t *testing.T) {
						tags := ctrls.TagsNonMatchingMissingRequiredExample()

						rejected := tagFilter.Reject(tags)

						t.Run("Rejected", func(t *testing.T) {
							Assert(t, rejected)
						})
					})

					t.Run("Missing required one of tags", func(t *testing.T) {
						tags := ctrls.TagsNonMatchingMissingAllRequiredOneOfExample()

						rejected := tagFilter.Reject(tags)

						t.Run("Rejected", func(t *testing.T) {
							Assert(t, rejected)
						})
					})
				})
			})
		})
	})
}
