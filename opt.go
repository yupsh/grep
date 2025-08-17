package command

type IgnoreCaseFlag bool

const (
	IgnoreCase    IgnoreCaseFlag = true
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

type flags struct {
	IgnoreCase   IgnoreCaseFlag
	LineNumber   LineNumberFlag
	Count        CountFlag
	Invert       InvertFlag
	WholeWord    WholeWordFlag
	FixedStrings FixedStringsFlag
	Recursive    RecursiveFlag
	FilesOnly    FilesOnlyFlag
	Quiet        QuietFlag
}

func (f IgnoreCaseFlag) Configure(flags *flags)   { flags.IgnoreCase = f }
func (f LineNumberFlag) Configure(flags *flags)   { flags.LineNumber = f }
func (f CountFlag) Configure(flags *flags)        { flags.Count = f }
func (f InvertFlag) Configure(flags *flags)       { flags.Invert = f }
func (f WholeWordFlag) Configure(flags *flags)    { flags.WholeWord = f }
func (f FixedStringsFlag) Configure(flags *flags) { flags.FixedStrings = f }
func (f RecursiveFlag) Configure(flags *flags)    { flags.Recursive = f }
func (f FilesOnlyFlag) Configure(flags *flags)    { flags.FilesOnly = f }
func (f QuietFlag) Configure(flags *flags)        { flags.Quiet = f }
