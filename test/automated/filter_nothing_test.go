package rag_test

import (
	"testing"

	"github.com/thehungry-dev/rag/pkg/ctrls"
)

func TestTagFilterNothingAlwaysSelects(t *testing.T) {
	tagFilter := ctrls.TagFilterNothingExample()
	tags := ctrls.TagsMatchingExample()

	selected := tagFilter.Select(tags)

	if !selected {
		t.Error()
	}
}
