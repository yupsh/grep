package command_test

import (
	"errors"
	"testing"

	"github.com/gloo-foo/testable/assertion"
	"github.com/gloo-foo/testable/run"
	command "github.com/yupsh/grep"
)

func TestGrep_Basic(t *testing.T) {
	result := run.Command(command.Grep("b")).
		WithStdinLines("a", "b", "c").Run()
	assertion.NoError(t, result.Err)
	assertion.Lines(t, result.Stdout, []string{"b"})
}

func TestGrep_InvertMatch(t *testing.T) {
	result := run.Command(command.Grep("b", command.Invert)).
		WithStdinLines("a", "b", "c").Run()
	assertion.NoError(t, result.Err)
	assertion.Count(t, result.Stdout, 2)
}

func TestGrep_IgnoreCase(t *testing.T) {
	result := run.Command(command.Grep("B", command.IgnoreCase)).
		WithStdinLines("a", "b", "c").Run()
	assertion.NoError(t, result.Err)
	assertion.Lines(t, result.Stdout, []string{"b"})
}

func TestGrep_Count(t *testing.T) {
	result := run.Command(command.Grep("a", command.Count)).
		WithStdinLines("a", "aa", "b").Run()
	assertion.NoError(t, result.Err)
	assertion.Count(t, result.Stdout, 1)
}

func TestGrep_LineNumber(t *testing.T) {
	result := run.Command(command.Grep("b", command.LineNumber)).
		WithStdinLines("a", "b", "c").Run()
	assertion.NoError(t, result.Err)
	assertion.Count(t, result.Stdout, 1)
}

func TestGrep_FixedStrings(t *testing.T) {
	result := run.Command(command.Grep(".", command.FixedStrings)).
		WithStdinLines("a.b", "abc").Run()
	assertion.NoError(t, result.Err)
	assertion.Lines(t, result.Stdout, []string{"a.b"})
}

func TestGrep_EmptyInput(t *testing.T) {
	result := run.Quick(command.Grep("a"))
	assertion.NoError(t, result.Err)
	assertion.Empty(t, result.Stdout)
}

func TestGrep_InputError(t *testing.T) {
	result := run.Command(command.Grep("a")).
		WithStdinError(errors.New("read failed")).Run()
	assertion.ErrorContains(t, result.Err, "read failed")
}

