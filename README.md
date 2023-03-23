# Rag

Provides a tag filtering language to filter tagged text.

Tag filtering allows building developer tools with easy filtering capabilities in mind.

## Usage

The core of `rag` is `TagFilter`, a data structure that can accept or reject a string slice called "input tags".
Through this README, `TagFilter` will refer to the string provided to `rag.Parse` to build a `TagFilter`.

There are 3 main functions in `rag`:
- `rag.Parse` which converts a string, such as `foo,_untagged` into a `TagFilter`
- `TagFilter.Select` which returns `true` if the filter decided that the input tags are accepted and as such, the data tagged with the input tags should be preserved and not filtered out
- `TagFilter.Reject`, opposite of `Select`

## Filtering Language

The TagFilter `_all` will ensure nothing is filtered out, that is `Select` always returns `true`.

If a TagFilter is not set to `_all`, it should contain a list of tags, separated by comma. If the input tags have at least one of the tags listed in TagFilter, `Select` will return `true`.
A tag can be prefixed with a _modifier_, which is either `+` or `-`.

A tag prefixed with `+` **must be present on the input tags**. If multiple tags are prefixed with `+` they must all be present in the input tags. This means that the input tags must include all tags prefixed with a `+` and **at least one unprefixed tag** (if any present in the TagFilter).

If a tag is prefixed with `-`, it must be absent from the input tags.
Every empty list of input tags is rejected unless the keyword `_untagged` is present in input tags.

## Usage

```go
package main

import (
	"github.com/thehungry-dev/rag"
)

func main() {
  filter := rag.Parse("foo,tag1,_untagged,+bar,+baz")

  var selected bool

  // selected is true
  selected = filter.Select([]string{"bar","foo","baz"})

  // selected is false
  selected = filter.Select([]string{"foo","baz"})
}
```

## Examples

### `foo,tag1,_untagged,+bar,+baz`

`Select` will return `true` if the input tags satisfy **all the following conditions**:

- Have at least one of: `foo` or `tag1`
- Have **both**: `bar` and `baz`

Alternatively the input tags could be empty.

### `foo,tag1,+bar,-baz`


`Select` will return `true` if the input tags satisfy **all the following conditions**:

- Have at least one of: `foo` or `tag1`
- Have: `bar`
- **Not have**: `baz`

## Development

### Testing

Requirements:

- `make`
