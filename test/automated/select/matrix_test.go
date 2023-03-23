package rag_test

import (
	_ "embed"
	"strings"
	"testing"

	"github.com/thehungry-dev/rag"
	"github.com/thehungry-dev/testmatrix"
)

//go:embed matrix.csv
var matrixFile string

func TestTagFilterSelectMatrix(t *testing.T) {
	matrix := testmatrix.NewColoredTestMatrix(t)
	reader := strings.NewReader(matrixFile)

	matrix.AssertCsv(reader, func(filterText string, tagsText string) bool {
		tags := strings.Split(tagsText, rag.TagSeparator)
		if tagsText == "" {
			tags = []string{}
		}

		filter := rag.Parse(filterText)

		return filter.Select(tags)
	})
}
