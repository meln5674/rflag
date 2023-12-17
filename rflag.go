package rflag

import (
	goflag "flag"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/huandu/xstrings"
	"github.com/pkg/errors"

	"github.com/spf13/pflag"

	goflagwrap "github.com/meln5674/rflag/pkg/goflag"
	pflagwrap "github.com/meln5674/rflag/pkg/pflag"
	"github.com/meln5674/rflag/pkg/rflag"
)

var (
	DefaultName func(string) string = xstrings.ToKebabCase
)

const (
	StructTag               = "rflag"
	StructTagFieldName      = "name"
	StructTagFieldShorthand = "shorthand"
	StructTagFieldUsage     = "usage"
	StructTagFieldPrefix    = "prefix"
	StructTagFieldSliceType = "slice-type"
	SliceTypeArray          = "array"
	SliceTypeSlice          = "slice"
	DefaultSliceType        = SliceTypeArray
)

type TagInfo = rflag.TagInfo
type FlagSet = rflag.FlagSet

func ParseTag(tag string) (TagInfo, error) {
	var info, empty TagInfo
	ix := 0
	for ix < len(tag) {
		nameStart := ix
		nameLen := strings.Index(tag[ix:], "=")
		if nameLen == -1 {
			return empty, fmt.Errorf("Expected name=value pair, got `%s`, did you forget to escape a comma?", tag[ix:])
		}
		ix += nameLen
		nameEnd := ix
		ix++

		valueStart := ix
		for ix < len(tag) {
			next := strings.Index(tag[ix:], ",")

			if next == -1 {
				ix = len(tag)
				break
			}
			ix += next
			if ix+1 == len(tag) || tag[ix+1] != ',' {
				break
			}
			ix += 2
		}
		valueEnd := ix
		ix++

		name := tag[nameStart:nameEnd]
		value := tag[valueStart:valueEnd]
		value = strings.ReplaceAll(value, ",,", ",")

		switch name {
		case StructTagFieldName:
			info.Name = value
		case StructTagFieldShorthand:
			info.Shorthand = value
		case StructTagFieldUsage:
			info.Usage = value
		case StructTagFieldPrefix:
			info.Prefix = value
		case StructTagFieldSliceType:
			switch value {
			case SliceTypeSlice, SliceTypeArray:
				info.SliceType = value
			default:
				return empty, fmt.Errorf("Unknown slice-type %s", value)
			}
		default:
			return empty, fmt.Errorf("Unknown struct tag field '%s'", name)
		}
	}
	if info.SliceType == "" {
		info.SliceType = DefaultSliceType
	}
	return info, nil
}

func registerFieldWithType[T any](f func(*T, T, TagInfo), ptrV reflect.Value, info TagInfo) {
	f(ptrV.Interface().(*T), ptrV.Elem().Interface().(T), info)
}

func registerField(flags FlagSet, prefix string, field reflect.StructField, ptrV reflect.Value) error {
	tag, ok := field.Tag.Lookup(StructTag)
	if !ok {
		return nil
	}

	info, err := ParseTag(tag)
	if err != nil {
		return fmt.Errorf("malformed rflag struct tag: %#v", err)
	}

	if info.Name == "" {
		info.Name = DefaultName(field.Name)
	}
	info.Name = prefix + info.Name

	switch ptrV.Elem().Interface().(type) {
	case bool:
		registerFieldWithType[bool](flags.BoolVar, ptrV, info)
	case int:
		registerFieldWithType[int](flags.IntVar, ptrV, info)
	case string:
		registerFieldWithType[string](flags.StringVar, ptrV, info)
	case map[string]string:
		registerFieldWithType[map[string]string](flags.StringToStringVar, ptrV, info)
	case time.Duration:
		registerFieldWithType[time.Duration](flags.DurationVar, ptrV, info)
	case []string:
		switch info.SliceType {
		case SliceTypeSlice:
			registerFieldWithType[[]string](flags.StringSliceVar, ptrV, info)
		case SliceTypeArray:
			registerFieldWithType[[]string](flags.StringArrayVar, ptrV, info)
		default:
			panic(fmt.Sprintf("BUG: Unknown slice type %s", info.SliceType))
		}
	default:
		switch field.Type.Kind() {
		case reflect.Struct:
			info, err := ParseTag(tag)
			if err != nil {
				return errors.Wrap(err, "Malformed rflag struct tag")
			}
			err = register(flags, prefix+info.Prefix, ptrV)
			if err != nil {
				return errors.Wrap(err, "Nested struct is invalid")
			}
		default:
			return fmt.Errorf("Unsupported flag type %#v", field.Type.Kind())
		}
	}

	return nil
}

func register(flags FlagSet, prefix string, v reflect.Value) error {
	typ := v.Type()
	if typ.Kind() != reflect.Pointer || typ.Elem().Kind() != reflect.Struct {
		return fmt.Errorf("Register must be called with a pointer to a struct, not %#v", v)
	}
	if v.IsNil() {
		return fmt.Errorf("Register cannot be called with a nil pointer")
	}

	destType := typ.Elem()
	dest := v.Elem()

	for ix := 0; ix < dest.NumField(); ix++ {
		field := destType.Field(ix)
		err := registerField(flags, prefix, field, dest.Field(ix).Addr())
		if err != nil {
			return errors.Wrapf(err, "Field %#v.%s is invalid", destType, field.Name)
		}
	}

	return nil
}

func Register(flags FlagSet, prefix string, v interface{}) error {
	return register(flags, prefix, reflect.ValueOf(v))
}

func MustRegister(flags FlagSet, prefix string, v interface{}) {
	err := Register(flags, prefix, v)
	if err != nil {
		panic(err)
	}
}

func ForGoFlag(flags *goflag.FlagSet) FlagSet {
	return &goflagwrap.FlagSet{GoFlagSet: flags}
}

func ForPFlag(flags *pflag.FlagSet) FlagSet {
	return &pflagwrap.FlagSet{PFlagSet: flags}
}
