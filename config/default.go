package config

import "os"

const (
	defaultDepth  = 5
	defaultLimit  = 20
	defaultUnits  = ""
	defaultSort   = "size_desc"
	defaultToFile = ""
	//DefaultToFile = "<no file>"
)

func (c *Config) setDefaultValues() {

	if c.Analyzer.Depth == nil {
		d := defaultDepth
		c.Analyzer.Depth = &d
	}
	if c.Analyzer.Path == nil {
		dir, _ := os.Getwd()
		c.Analyzer.Path = &dir
	}

	if c.Printer.Limit == nil {
		l := defaultLimit
		c.Printer.Limit = &l
	}

	if c.Printer.Units == nil {
		u := defaultUnits
		c.Printer.Units = &u
	}

	if c.Printer.ToFile == nil {
		u := defaultToFile
		c.Printer.ToFile = &u
	}

	if c.Printer.Sort == nil {
		s := defaultSort
		c.Printer.Sort = &s
	}
}
