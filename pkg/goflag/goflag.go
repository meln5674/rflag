package goflag

import (
	"flag"
	"time"

	"github.com/meln5674/rflag/pkg/rflag"
)

type FlagSet struct {
	GoFlagSet *flag.FlagSet
}

func (f FlagSet) BoolVar(p *bool, value bool, info rflag.TagInfo) {
	f.GoFlagSet.BoolVar(p, info.Name, value, info.Usage)
}
func (f FlagSet) IntVar(p *int, value int, info rflag.TagInfo) {
	f.GoFlagSet.IntVar(p, info.Name, value, info.Usage)
}
func (f FlagSet) StringVar(p *string, value string, info rflag.TagInfo) {
	f.GoFlagSet.StringVar(p, info.Name, value, info.Usage)
}
func (f FlagSet) StringToStringVar(p *map[string]string, value map[string]string, info rflag.TagInfo) {
	panic("flag library does not support map variables")
}
func (f FlagSet) DurationVar(p *time.Duration, value time.Duration, info rflag.TagInfo) {
	f.GoFlagSet.DurationVar(p, info.Name, value, info.Usage)
}
func (f FlagSet) StringArrayVar(p *[]string, value []string, info rflag.TagInfo) {
	panic("flag library does not support slice variables")
}
func (f FlagSet) StringSliceVar(p *[]string, value []string, info rflag.TagInfo) {
	panic("flag library does not support slice variables")
}
