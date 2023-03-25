// Package ctrls provides building blocks to generate sample data for a test environment
package ctrls

import (
	"fmt"

	"github.com/thehungry-dev/rag"
)

func TagNameOneOfExample() string    { return "aTag" }
func TagNameRequiredExample() string { return "requiredTag" }
func TagNameExcludedExample() string { return "excludedTag" }

// aTag
func OneOfTagExample() string { return TagNameOneOfExample() }

// +requiredTag
func RequiredTagExample() string { return fmt.Sprintf("+%s", TagNameRequiredExample()) }

// -excludedTag
func ExcludeTagExample() string { return fmt.Sprintf("-%s", TagNameExcludedExample()) }

// aTag,requiredTag
func TagsMatchingExample() []string { return []string{TagNameOneOfExample(), TagNameRequiredExample()} }

// aTag,requiredTag,excludedTag
func TagsExcludedMatchingExample() []string {
	return []string{TagNameOneOfExample(), TagNameRequiredExample(), TagNameExcludedExample()}
}

// aTag
func TagsNonMatchingMissingRequiredExample() []string { return []string{TagNameOneOfExample()} }

// requiredTag
func TagsNonMatchingMissingAllRequiredOneOfExample() []string {
	return []string{TagNameRequiredExample()}
}

// aTag,+requiredTag,-excludedTag
func StringExample() string {
	return fmt.Sprintf("%s,%s,%s", OneOfTagExample(), RequiredTagExample(), ExcludeTagExample())
}

// +requiredTag,-excludedTag
func StringExcludedRequiredExample() string {
	return fmt.Sprintf("%s,%s", RequiredTagExample(), ExcludeTagExample())
}

func TagFilterExample() *rag.TagFilter {
	return rag.Parse(StringExample())
}
func TagFilterRequiredExample() *rag.TagFilter {
	return rag.Parse(StringExcludedRequiredExample())
}
func TagFilterNothingExample() *rag.TagFilter {
	return rag.FilterNothing
}
