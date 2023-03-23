// Package rag provides a tag filtering language for tagged text
package rag

import (
	"strings"

	"github.com/thehungry-dev/rag/internal/set"
)

const TagSeparator = ","
const TagRequiredToken = "+"
const TagExcludedToken = "-"
const AllTagsKeyword = "_all"
const UntaggedKeyword = "_untagged"

// Filter for which Select always returns true
var FilterNothing *TagFilter

type TagFilter struct {
	requiredTags set.Set
	oneOfTags    set.Set
	excludedTags set.Set
	untagged     bool
}

// Converts rag language into a filtering object
func Parse(filterText string) *TagFilter {
	text := strings.TrimSpace(filterText)

	if text == "" || text == AllTagsKeyword {
		return FilterNothing
	}

	tags := strings.Split(text, TagSeparator)
	maxTags := len(tags)

	requiredTags := set.BuildSetOfSize(maxTags)
	oneOfTags := set.BuildSetOfSize(maxTags)
	excludedTags := set.BuildSetOfSize(maxTags)
	filter := &TagFilter{requiredTags, oneOfTags, excludedTags, false}

	for _, tag := range tags {
		tag = strings.TrimSpace(tag)

		if tag == "" {
			continue
		}

		switch {
		case tag == UntaggedKeyword:
			filter.untagged = true
		case strings.HasPrefix(tag, TagRequiredToken):
			trimmedTag := strings.TrimPrefix(tag, TagRequiredToken)
			requiredTags.Add(trimmedTag)
		case strings.HasPrefix(tag, TagExcludedToken):
			trimmedTag := strings.TrimPrefix(tag, TagExcludedToken)
			excludedTags.Add(trimmedTag)
		default:
			oneOfTags.Add(tag)
		}
	}

	return filter
}

func (tagFilter *TagFilter) IsRequiredTag(tag string) bool {
	if tagFilter == nil {
		return false
	}
	return tagFilter.requiredTags.Include(tag)
}
func (tagFilter *TagFilter) IsExcludedTag(tag string) bool {
	if tagFilter == nil {
		return false
	}
	return tagFilter.excludedTags.Include(tag)
}
func (tagFilter *TagFilter) IsRequiredOneOfTag(tag string) bool {
	if tagFilter == nil {
		return false
	}
	return tagFilter.oneOfTags.Include(tag)
}
func (tagFilter *TagFilter) String() string {
	if tagFilter == nil || tagFilter == FilterNothing {
		return "_all"
	}

	tags := make([]string, 0)
	if tagFilter.untagged {
		tags = append(tags, UntaggedKeyword)
	}

	for tag := range tagFilter.oneOfTags {
		tags = append(tags, tag)
	}
	for tag := range tagFilter.requiredTags {
		tags = append(tags, "+"+tag)
	}
	for tag := range tagFilter.excludedTags {
		tags = append(tags, "-"+tag)
	}

	return strings.Join(tags, TagSeparator)
}
func (tagFilter *TagFilter) IsUntaggedOnly() bool {
	return tagFilter.oneOfTags.IsEmpty() &&
		tagFilter.excludedTags.IsEmpty() &&
		tagFilter.requiredTags.IsEmpty() &&
		tagFilter.untagged
}

func (tagFilter *TagFilter) Select(inputTags []string) bool {
	if tagFilter == nil || tagFilter == FilterNothing {
		return true
	}

	if len(inputTags) == 0 {
		return tagFilter.untagged
	}

	if !tagFilter.requiredTags.SubsetOf(inputTags) {
		return false
	}

	if !tagFilter.excludedTags.IsEmpty() && tagFilter.excludedTags.IncludeAny(inputTags) {
		return false
	}

	if !tagFilter.oneOfTags.IncludeAny(inputTags) {
		return false
	}

	if tagFilter.IsUntaggedOnly() && len(inputTags) > 0 {
		return false
	}

	return true
}

func (tagFilter *TagFilter) Reject(inputTags []string) bool {
	return !tagFilter.Select(inputTags)
}
