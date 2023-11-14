package rflag

import "time"

type TagInfo struct {
	Name      string
	Shorthand string
	Usage     string
	Prefix    string
}

type FlagSet interface {
	BoolVar(p *bool, value bool, info TagInfo)
	IntVar(p *int, value int, info TagInfo)
	StringVar(p *string, value string, info TagInfo)
	StringToStringVar(p *map[string]string, value map[string]string, info TagInfo)
	DurationVar(p *time.Duration, value time.Duration, info TagInfo)
}
