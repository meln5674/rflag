package goflag

import (
	"time"

	flag "github.com/spf13/pflag"

	"github.com/meln5674/rflag/pkg/rflag"
)

type FlagSet struct {
	PFlagSet *flag.FlagSet
}

func (f FlagSet) BoolVar(p *bool, value bool, info rflag.TagInfo) {
	f.PFlagSet.BoolVar(p, info.Name, value, info.Usage)
}
func (f FlagSet) IntVar(p *int, value int, info rflag.TagInfo) {
	f.PFlagSet.IntVar(p, info.Name, value, info.Usage)
}
func (f FlagSet) StringVar(p *string, value string, info rflag.TagInfo) {
	f.PFlagSet.StringVar(p, info.Name, value, info.Usage)
}
func (f FlagSet) StringToStringVar(p *map[string]string, value map[string]string, info rflag.TagInfo) {
	f.PFlagSet.StringToStringVar(p, info.Name, value, info.Usage)
}
func (f FlagSet) DurationVar(p *time.Duration, value time.Duration, info rflag.TagInfo) {
	f.PFlagSet.DurationVar(p, info.Name, value, info.Usage)
}
