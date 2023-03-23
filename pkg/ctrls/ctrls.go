// Package ctrls provides building blocks to generate sample data for a test environment
package ctrls

import (
	"fmt"

	"github.com/thehungry-dev/rag"
)

func TagNameOneOfExample() string    { return "aTag" }
func TagNameRequiredExample() string { return "requiredTag" }
func TagNameExcludedExample() string { return "excludedTag" }

func OneOfTagExample() string    { return TagNameOneOfExample() }
func RequiredTagExample() string { return fmt.Sprintf("+%s", TagNameRequiredExample()) }
func ExcludeTagExample() string  { return fmt.Sprintf("-%s", TagNameExcludedExample()) }

func TagsMatchingExample() []string { return []string{TagNameOneOfExample(), TagNameRequiredExample()} }
func TagsExcludedMatchingExample() []string {
	return []string{TagNameOneOfExample(), TagNameRequiredExample(), TagNameExcludedExample()}
}

func TagsNonMatchingMissingRequiredExample() []string { return []string{TagNameOneOfExample()} }
func TagsNonMatchingMissingAllRequiredOneOfExample() []string {
	return []string{TagNameRequiredExample()}
}

func StringExample() string {
	return fmt.Sprintf("%s,%s,%s", OneOfTagExample(), RequiredTagExample(), ExcludeTagExample())
}
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
