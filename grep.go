package grep

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"os"
	"regexp"

	yup "github.com/yupsh/framework"
	"github.com/yupsh/framework/opt"
	localopt "github.com/yupsh/grep/opt"
)

// Flags represents the configuration options for the grep command
type Flags = localopt.Flags

// Command implementation
type command opt.Inputs[string, Flags]

// Grep creates a new grep command with the given parameters
// First parameter should be the pattern, rest are files
func Grep(parameters ...any) yup.Command {
	return command(opt.Args[string, Flags](parameters...))
}

func (c command) Execute(ctx context.Context, stdin io.Reader, stdout, stderr io.Writer) error {
	if len(c.Positional) == 0 {
		fmt.Fprintln(stderr, "grep: missing pattern")
		return fmt.Errorf("missing pattern")
	}

	pattern := c.Positional[0]
	files := c.Positional[1:]

	// Compile regex pattern
	var regex *regexp.Regexp
	var err error

	if c.Flags.FixedStrings {
		// Escape special regex characters for literal matching
		pattern = regexp.QuoteMeta(pattern)
	}

	if c.Flags.WholeWord {
		pattern = `\b` + pattern + `\b`
	}

	flags := ""
	if c.Flags.IgnoreCase {
		flags = "(?i)"
	}

	regex, err = regexp.Compile(flags + pattern)
	if err != nil {
		fmt.Fprintf(stderr, "grep: invalid pattern: %v\n", err)
		return err
	}

	// If no files specified, read from stdin
	if len(files) == 0 {
		totalMatches := c.searchReader(ctx, stdin, stdout, "", regex)
		if totalMatches == 0 {
			return fmt.Errorf("no matches found")
		}
		return nil
	}

	// Process each file
	totalMatches := 0
	for _, filename := range files {
		// Check for cancellation before each file
		if err := yup.CheckContextCancellation(ctx); err != nil {
			return err
		}

		file, err := os.Open(filename)
		if err != nil {
			fmt.Fprintf(stderr, "grep: %s: %v\n", filename, err)
			continue
		}

		matches := c.searchReader(ctx, file, stdout, filename, regex)
		totalMatches += matches
		file.Close()
	}

	// Return non-zero exit if no matches found (grep convention)
	if totalMatches == 0 {
		return fmt.Errorf("no matches found")
	}

	return nil
}

func (c command) searchReader(ctx context.Context, reader io.Reader, output io.Writer, filename string, regex *regexp.Regexp) int {
	scanner := bufio.NewScanner(reader)
	lineNum := 0
	matchCount := 0

	for yup.ScanWithContext(ctx, scanner) {
		lineNum++
		line := scanner.Text()

		matches := regex.MatchString(line)

		// Apply invert logic
		if c.Flags.Invert {
			matches = !matches
		}

		if matches {
			matchCount++

			// Skip output if quiet mode
			if c.Flags.Quiet {
				continue
			}

			// Handle different output modes
			if c.Flags.Count {
				// Count mode - will print total at end
				continue
			}

			if c.Flags.FilesOnly {
				if filename != "" {
					fmt.Fprintln(output, filename)
				}
				break // Only need to find one match per file
			}

			// Normal output mode
			prefix := ""
			if filename != "" && len(c.Positional) > 2 { // Multiple files
				prefix = filename + ":"
			}

			if c.Flags.LineNumber {
				prefix += fmt.Sprintf("%d:", lineNum)
			}

			fmt.Fprintf(output, "%s%s\n", prefix, line)
		}
	}

	// Check if context was cancelled
	if err := yup.CheckContextCancellation(ctx); err != nil {
		return matchCount // Return partial results
	}

	// Print count if in count mode
	if bool(c.Flags.Count) && !bool(c.Flags.Quiet) {
		prefix := ""
		if filename != "" && len(c.Positional) > 2 {
			prefix = filename + ":"
		}
		fmt.Fprintf(output, "%s%d\n", prefix, matchCount)
	}

	return matchCount
}
