package grep_test

import (
	"context"
	"os"
	"strings"

	"github.com/yupsh/grep"
	"github.com/yupsh/grep/opt"
)

func ExampleGrep() {
	ctx := context.Background()
	input := strings.NewReader("error occurred\ninfo message\nerror found\nwarning issued\n")

	cmd := grep.Grep("error")
	cmd.Execute(ctx, input, os.Stdout, os.Stderr)
	// Output:
	// error occurred
	// error found
}

func ExampleGrep_withLineNumbers() {
	ctx := context.Background()
	input := strings.NewReader("first line\nerror here\nthird line\nerror again\n")

	cmd := grep.Grep("error", opt.LineNumber)
	cmd.Execute(ctx, input, os.Stdout, os.Stderr)
	// Output:
	// 2:error here
	// 4:error again
}

func ExampleGrep_caseInsensitive() {
	ctx := context.Background()
	input := strings.NewReader("Error occurred\nINFO message\nerror found\nWARNING issued\n")

	cmd := grep.Grep("error", opt.IgnoreCase, opt.LineNumber)
	cmd.Execute(ctx, input, os.Stdout, os.Stderr)
	// Output:
	// 1:Error occurred
	// 3:error found
}

func ExampleGrep_countOnly() {
	ctx := context.Background()
	input := strings.NewReader("error 1\ninfo\nerror 2\nwarning\nerror 3\n")

	cmd := grep.Grep("error", opt.Count)
	cmd.Execute(ctx, input, os.Stdout, os.Stderr)
	// Output:
	// 3
}

func ExampleGrep_invert() {
	ctx := context.Background()
	input := strings.NewReader("error line\ngood line\nerror again\nanother good\n")

	cmd := grep.Grep("error", opt.Invert)
	cmd.Execute(ctx, input, os.Stdout, os.Stderr)
	// Output:
	// good line
	// another good
}

func ExampleGrep_wholeWord() {
	ctx := context.Background()
	input := strings.NewReader("error\nerrors\nno error here\nan error occurred\n")

	cmd := grep.Grep("error", opt.WholeWord)
	cmd.Execute(ctx, input, os.Stdout, os.Stderr)
	// Output:
	// error
	// no error here
	// an error occurred
}
