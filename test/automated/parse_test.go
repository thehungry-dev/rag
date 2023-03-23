package filter_test

import (
	"testing"

	ctrls "github.com/thehungry-dev/rag/ctrls/tag/filter"
	"github.com/thehungry-dev/rag/tag/filter"
)

func TestTagFilterParse(t *testing.T) {
	t.Run("Tag", func(t *testing.T) {
		t.Run("Filter", func(t *testing.T) {
			t.Run("Parse", func(t *testing.T) {
				tagFilterString := ctrls.StringExample()

				t.Run("One of tags", func(t *testing.T) {
					tagFilter := filter.Parse(tagFilterString)
					oneOfTagName := ctrls.TagNameOneOfExample()

					t.Run("Included", func(t *testing.T) {
						Assert(t, tagFilter.IsRequiredOneOfTag(oneOfTagName))
					})
				})

				t.Run("Required tags", func(t *testing.T) {
					tagFilter := filter.Parse(tagFilterString)
					requiredTagName := ctrls.TagNameRequiredExample()

					t.Run("Included", func(t *testing.T) {
						Assert(t, tagFilter.IsRequiredTag(requiredTagName))
					})
				})

				t.Run("Excluded tags", func(t *testing.T) {
					tagFilter := filter.Parse(tagFilterString)
					excludedTagName := ctrls.TagNameExcludedExample()

					t.Run("Included", func(t *testing.T) {
						Assert(t, tagFilter.IsExcludedTag(excludedTagName))
					})
				})
			})
		})
	})
}
