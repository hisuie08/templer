package option

type Option struct {
	DataArgs   []string
	DataFormat string
	OutArg     string
	SetValues  []string
	LoadEnv    bool
	Template   Template
}

type Template struct {
	Value     string
	AsLiteral bool
	Suffix    string
}

const OutSibling = "__SIBLING__"

const MetaStr = "*?[]"

var Prefix = struct {
	Str  string
	File string
	Glob string
}{Str: "str:", File: "file:", Glob: "glob:"}
