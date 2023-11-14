# RFlag - Reflection-based Command Line Parsing

This Library implements a reflection-based command line parser that can use either the standard library "flag" package or [pflag](github.com/spf13/pflag).

For usage, see the following example

```go
import (
    goflag "flag"

    "github.com/spf13/plfag"
)

// Define any number of structs containing your command line flags, and tag its fields with the rflag struct tag
// The value of this tag is a nested set of name=value pairs, separated by comma
// Currently, string, int, bool, and map[string]string fields are supported, as well as structs (including embedded)
type Args struct {
    Foo string `rflag:"usage=This is the help message for your flag"` 
    Bar int `rflag:"name=something-else,usage=This flag has a specific name instead of the field name in kebab-case"` 
    Baz bool `rflag:"shorthand=b,usage=This flag has a shorthand flag -b"` 
    Qux struct{
       Quux map[string]string `rflag:"name=quux,usage=This flag is in a nested struct field with a prefix,, and the usage message contains escaped commas using a double comma"` 
    } `rflag:"prefix=nested-"`
}



func main() {
    // Next, declare a variable with the default values
    // This can also run in init() with global variables instead
    var args = Args{
        Foo: "",
        Bar: 5,
        // Baz: false, // Zero values are acceptable
        Qux: map[string]string{"default": "value"},
    }

    // Finally, register and parse your arguments

    goFlagSet := goflag.CommandLine() // Or use your own flag set
    err := rflag.Register(rflag.ForGoFlag(goFlagSet, "", &args)) // The empty string is an optional prefix to apply to all flags
    // To panic instead of returning error, call MustRegister instead
    goflag.Parse() // or call goFlagset.Parse()

    // Or

    pflagSet := pflag.CommandLine() // Or use your own flag set, or the flag set of a cobra command
    err := rflag.Register(rflag.ForPFlag(pflagSet, &args))
    pflag.Parse() // or call pflagSet.Parse(), or use cobra

    // At this point, if the arguments provided were
    // --foo=value1 -b --nested-quux x=y --nested-quux a=b 
    // args would then be
    // Args{Foo: "value1", Bar: 5, Baz: true, Qux: struct{Quux map[string]string}{Quux: map[string]string{"x": "y", "a": "b"}}
}
```

This library is compatible with manually assigning flags, though take care to avoid duplicate flag definitions.
