package filter_test

import (
	_ "embed"
	"strings"
	"testing"

	"github.com/thehungry-dev/rag/tag/filter"
	"github.com/thehungry-dev/testmatrix"
)

//go:embed matrix.csv
var matrixFile string

func TestTagFilterSelectMatrix(t *testing.T) {
	t.Run("Tag", func(t *testing.T) {
		t.Run("Filter", func(t *testing.T) {
			t.Run("Select", func(t *testing.T) {
				t.Run("Matrix", func(t *testing.T) {
					matrix := testmatrix.NewColoredTestMatrix(t)
					reader := strings.NewReader(matrixFile)

					matrix.AssertCsv(reader, func(filterText string, tagsText string) bool {
						tags := strings.Split(tagsText, filter.TagSeparator)
						if tagsText == "" {
							tags = []string{}
						}

						filter := filter.Parse(filterText)

						return filter.Select(tags)
					})
				})
			})
		})
	})
}
