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
