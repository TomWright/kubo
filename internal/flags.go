package internal

import (
	"fmt"
	"strings"
)

// OverrideFlag allows us to collect a list of overrides from commandline flags.
type OverrideFlag []*Override

// Type returns the type of flag.
// As far as I can tell this isn't used...
// todo : figure out what this is supposed to return.
func (of *OverrideFlag) Type() string {
	panic("implement me")
}

// Override represents a single override.
type Override struct {
	// Path is a string that may have dot separators within it.
	Path string
	// Value is the value to set in the config file.
	Value string
}

// String returns a string representation of the override.
func (o Override) String() string {
	return fmt.Sprintf("%s=%s", o.Path, o.Value)
}

// String returns a string representation of all Override's within the OverrideFlag.
func (of *OverrideFlag) String() string {
	val := make([]string, len(*of))
	for k, o := range *of {
		val[k] = o.String()
	}
	return strings.Join(val, " ")
}

// Set is used to add a new value to the OverrideFlag.
func (of *OverrideFlag) Set(value string) error {
	args := strings.Split(value, "=")
	switch len(args) {
	case 0:
		// No value was given.
		return nil
	case 1:
		// A blank value was given.
		*of = append(*of, &Override{
			Path:  args[0],
			Value: "",
		})
	case 2:
		// A single value was given.
		*of = append(*of, &Override{
			Path:  args[0],
			Value: args[1],
		})
	default:
		// The value contained more than 1 = sign.
		// Assume the extra = values are within the value.
		*of = append(*of, &Override{
			Path:  args[0],
			Value: strings.Join(args[1:], "="),
		})
	}
	return nil
}
