# yup.grep Examples

This directory contains runnable Go examples that demonstrate how to use the `yup.grep` command.

## Running Examples

The examples are implemented as Go example functions in the test files:

```bash
# Run all examples
go test -v

# Run specific examples
go test -run ExampleGrep
go test -run ExampleGrep_caseInsensitive

# View examples in documentation
go doc -all
```

## Available Examples

See `../grep_test.go` for the actual runnable examples:

- `ExampleGrep()` - Basic pattern searching
- `ExampleGrep_withLineNumbers()` - Search with line numbers
- `ExampleGrep_caseInsensitive()` - Case insensitive matching
- `ExampleGrep_countOnly()` - Count matches only
- `ExampleGrep_invert()` - Show non-matching lines
- `ExampleGrep_wholeWord()` - Match whole words only

## Usage Patterns

### Basic Search
```go
import "github.com/nicerobot/yup.grep"
import "github.com/nicerobot/yup.grep/flags"

cmd := grep.Grep("pattern", "file.txt", flags.IgnoreCase, flags.LineNumber)
err := cmd.Execute(ctx, input, output, stderr)
```

### Pipeline Integration
```go
pipeline := yup.Pipe(
    cat.Cat("app.log"),
    grep.Grep("ERROR", flags.IgnoreCase, flags.LineNumber),
    head.Head(flags.LineCount(10)),
)
```

## Available Flags

- `flags.IgnoreCase` / `flags.CaseSensitive` - Case sensitivity
- `flags.LineNumber` / `flags.NoLineNumber` - Show line numbers
- `flags.Count` / `flags.NoCount` - Count matches only
- `flags.Invert` / `flags.NoInvert` - Invert matching
- `flags.WholeWord` / `flags.NoWholeWord` - Match whole words
- `flags.FixedStrings` / `flags.NoFixedStrings` - Literal string matching
- `flags.FilesOnly` / `flags.NoFilesOnly` - Show filenames only
- `flags.Quiet` / `flags.NoQuiet` - Suppress output