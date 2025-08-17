package opt

// Boolean flag types with constants
type IgnoreCaseFlag bool
const (
	IgnoreCase   IgnoreCaseFlag = true
	CaseSensitive IgnoreCaseFlag = false
)

type LineNumberFlag bool
const (
	LineNumber   LineNumberFlag = true
	NoLineNumber LineNumberFlag = false
)

type CountFlag bool
const (
	Count   CountFlag = true
	NoCount CountFlag = false
)

type InvertFlag bool
const (
	Invert   InvertFlag = true
	NoInvert InvertFlag = false
)

type WholeWordFlag bool
const (
	WholeWord   WholeWordFlag = true
	NoWholeWord WholeWordFlag = false
)

type FixedStringsFlag bool
const (
	FixedStrings   FixedStringsFlag = true
	NoFixedStrings FixedStringsFlag = false
)

type RecursiveFlag bool
const (
	Recursive   RecursiveFlag = true
	NoRecursive RecursiveFlag = false
)

type FilesOnlyFlag bool
const (
	FilesOnly   FilesOnlyFlag = true
	NoFilesOnly FilesOnlyFlag = false
)

type QuietFlag bool
const (
	Quiet   QuietFlag = true
	NoQuiet QuietFlag = false
)

// Flags represents the configuration options for the grep command
type Flags struct {
	IgnoreCase   IgnoreCaseFlag   // Case insensitive matching
	LineNumber   LineNumberFlag   // Show line numbers
	Count        CountFlag        // Only show count of matches
	Invert       InvertFlag       // Invert match (show non-matching lines)
	WholeWord    WholeWordFlag    // Match whole words only
	FixedStrings FixedStringsFlag // Treat pattern as literal string
	Recursive    RecursiveFlag    // Search recursively
	FilesOnly    FilesOnlyFlag    // Only show filenames with matches
	Quiet        QuietFlag        // Suppress output, only return exit status
}

// Flag configuration methods
func (f IgnoreCaseFlag) Configure(flags *Flags) { flags.IgnoreCase = f }
func (f LineNumberFlag) Configure(flags *Flags) { flags.LineNumber = f }
func (f CountFlag) Configure(flags *Flags) { flags.Count = f }
func (f InvertFlag) Configure(flags *Flags) { flags.Invert = f }
func (f WholeWordFlag) Configure(flags *Flags) { flags.WholeWord = f }
func (f FixedStringsFlag) Configure(flags *Flags) { flags.FixedStrings = f }
func (f RecursiveFlag) Configure(flags *Flags) { flags.Recursive = f }
func (f FilesOnlyFlag) Configure(flags *Flags) { flags.FilesOnly = f }
func (f QuietFlag) Configure(flags *Flags) { flags.Quiet = f }
