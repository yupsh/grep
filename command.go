package command

import (
	"fmt"
	"regexp"
	"strings"

	yup "github.com/gloo-foo/framework"
)

type Pattern string

// grepCommand stores both pattern and file inputs
type grepCommand struct {
	pattern Pattern
	files   yup.Inputs[yup.File, flags]
}

func Grep(pattern Pattern, parameters ...any) yup.Command {
	// Pattern is separate, remaining parameters are files
	files := yup.Initialize[yup.File, flags](parameters...)
	return grepCommand{
		pattern: pattern,
		files:   files,
	}
}

func (p grepCommand) Executor() yup.CommandExecutor {
	pattern := string(p.pattern)

	// Compile regex if not fixed strings
	var re *regexp.Regexp
	var compileErr error

	if !bool(p.files.Flags.FixedStrings) {
		flags := ""
		if bool(p.files.Flags.IgnoreCase) {
			flags = "(?i)"
		}
		re, compileErr = regexp.Compile(flags + pattern)
	}

	return p.files.Wrap(
		yup.StatefulLineTransform(func(lineNum int64, line string) (string, bool) {
		if compileErr != nil && !bool(p.files.Flags.FixedStrings) {
			return "", false
		}

		// Check if line matches
		matched := false

		if bool(p.files.Flags.FixedStrings) {
			// Fixed string matching
			searchIn := line
			searchFor := pattern
			if bool(p.files.Flags.IgnoreCase) {
				searchIn = strings.ToLower(line)
				searchFor = strings.ToLower(pattern)
			}

			if bool(p.files.Flags.WholeWord) {
				// Simple word boundary check
				words := strings.Fields(searchIn)
				for _, word := range words {
					if word == searchFor {
						matched = true
						break
					}
				}
			} else {
				matched = strings.Contains(searchIn, searchFor)
			}
		} else {
			// Regex matching
			matched = re.MatchString(line)
		}

		// Invert match if flag is set
		if bool(p.files.Flags.Invert) {
			matched = !matched
		}

		// Return based on flags
		if bool(p.files.Flags.Quiet) {
			// Quiet mode - no output
			return "", false
		}

		if !matched {
			return "", false
		}

		// Format output
		output := line
		if bool(p.files.Flags.LineNumber) {
			output = fmt.Sprintf("%d:%s", lineNum, line)
		}

		return output, true
	}).Executor(),
	)
}
