# yup.grep

A pure Go implementation of the `grep` command that adheres to the yup.sh Command interface.

## Features

- Pattern matching with regular expressions
- Case-sensitive and case-insensitive search
- Line numbering
- Count matches
- Invert matching
- Whole word matching
- Fixed string matching
- Multiple file support
- Strongly typed flag system

## Usage

```go
import "github.com/yupsh/yup.grep"

// Basic usage
cmd := grep.Grep("pattern", "file.txt")

// With flags
cmd := grep.Grep("error", "log.txt", grep.IgnoreCase{}, grep.LineNumber{})

// Search multiple files
cmd := grep.Grep("TODO", "file1.go", "file2.go", grep.Count{})

// Execute
err := cmd.Execute(ctx, input, output, stderr)
```

## Flags

- `IgnoreCase{}` - Case insensitive matching (-i)
- `LineNumber{}` - Show line numbers (-n)
- `Count{}` - Only show count of matches (-c)
- `Invert{}` - Invert match, show non-matching lines (-v)
- `WholeWord{}` - Match whole words only (-w)
- `FixedStrings{}` - Treat pattern as literal string (-F)
- `Recursive{}` - Search recursively (-r)
- `FilesOnly{}` - Only show filenames with matches (-l)
- `Quiet{}` - Suppress output, only return exit status (-q)
