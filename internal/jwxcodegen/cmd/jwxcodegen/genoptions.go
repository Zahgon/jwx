package main

type OptionsObjects struct {
	Output      string
	PackageName string   `yaml:"package_name"`
	Imports     []string `yaml:"imports"`
	Interfaces  []*struct {
		Name         string
		Comment      string
		ConcreteType string `yaml:"concrete_type"`
		Methods      []string
		Embeds       []string
	} `yaml:"interfaces"`
	Options []*struct {
		Ident         string
		OptionName    string `yaml:"option_name"` // usually "With" + $Ident
		IdentName     string `yaml:"ident_name"`
		SkipOption    bool   `yaml:"skip_option"`
		Interface     string
		ConcreteType  string
		Comment       string
		ArgumentType  string `yaml:"argument_type"`
		ConstantValue string `yaml:"constant_value"`
	} `yaml:"options"`
}

func runOptions(args []string) error { _ = "STUB: not implemented"; return nil }

func genOptionsFile(objects *OptionsObjects) error { _ = "STUB: not implemented"; return nil }

// for some reason without this the goimports in my environment tries to import a differnet package

// Write all imports -- they will be pruned by golang.org/x/tools/imports eventually,
// so it's okay to be redundant

func genOptionTests(objects *OptionsObjects) error { _ = "STUB: not implemented"; return nil }

func runAllOptions(_ []string) error { _ = "STUB: not implemented"; return nil }
