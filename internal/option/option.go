package option

type Option struct {
	DataArgs            []string
	DataStrictJson      bool
	OutDir              string
	OutDefault          bool
	AllowShellExecution bool
	AllowedShell        []string
	SetValues           []string
	AllowEnv            bool
	Template            Template
}

type Template struct {
	Value     string
	AsLiteral bool
	Suffix    string
}

const MetaStr = "*?[]"

var Prefix = struct {
	Str  string
	File string
	Glob string
}{Str: "str:", File: "file:", Glob: "glob:"}
