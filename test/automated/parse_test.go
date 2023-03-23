package rag_test

import (
	"testing"

	"github.com/thehungry-dev/rag"
	"github.com/thehungry-dev/rag/pkg/ctrls"
)

func TestParse(t *testing.T) {
	t.Parallel()
	tagFilterString := ctrls.StringExample()

	t.Run("One of tags included", func(t *testing.T) {
		tagFilter := rag.Parse(tagFilterString)
		oneOfTagName := ctrls.TagNameOneOfExample()

		included := tagFilter.IsRequiredOneOfTag(oneOfTagName)

		if !included {
			t.Error()
		}
	})

	t.Run("Required tags included", func(t *testing.T) {
		tagFilter := rag.Parse(tagFilterString)
		requiredTagName := ctrls.TagNameRequiredExample()

		included := tagFilter.IsRequiredTag(requiredTagName)

		if !included {
			t.Error()
		}
	})

	t.Run("Excluded tags included", func(t *testing.T) {
		tagFilter := rag.Parse(tagFilterString)
		excludedTagName := ctrls.TagNameExcludedExample()

		excluded := tagFilter.IsExcludedTag(excludedTagName)

		if !excluded {
			t.Error()
		}
	})
}
