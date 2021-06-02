package cmd

import (
	"couture/internal/pkg/sink/pretty/theme"
	"github.com/alecthomas/kong"
	"github.com/araddon/dateparse"
	errors2 "github.com/pkg/errors"
	"reflect"
	"regexp"
	"strings"
	"time"
)

type (
	autoResize       bool
	banner           bool
	color            bool
	columns          []string
	consistentColors bool
	expandJSON       bool
	highlight        bool
	multiline        bool
	themeName        string
	timeFormat       string
	tty              bool
	width            uint
	wrap             bool
)

// AfterApply ...
//goland:noinspection GoUnnecessarilyExportedIdentifiers
func (v autoResize) AfterApply() error { prettyConfig.AutoResize = bool(v); return nil }

// AfterApply ...
//goland:noinspection GoUnnecessarilyExportedIdentifiers
func (v banner) AfterApply() error { prettyConfig.Banner = bool(v); return nil }

// AfterApply ...
//goland:noinspection GoUnnecessarilyExportedIdentifiers
func (v color) AfterApply() error { prettyConfig.Color = bool(v); return nil }

// AfterApply ...
//goland:noinspection GoUnnecessarilyExportedIdentifiers
func (v columns) AfterApply() error { prettyConfig.Columns = v; return nil }

// AfterApply ...
//goland:noinspection GoUnnecessarilyExportedIdentifiers
func (v consistentColors) AfterApply() error { prettyConfig.ConsistentColors = bool(v); return nil }

// AfterApply ...
//goland:noinspection GoUnnecessarilyExportedIdentifiers
func (v expandJSON) AfterApply() error { prettyConfig.ExpandJSON = bool(v); return nil }

// AfterApply ...
//goland:noinspection GoUnnecessarilyExportedIdentifiers
func (v highlight) AfterApply() error { prettyConfig.Highlight = bool(v); return nil }

// AfterApply ...
//goland:noinspection GoUnnecessarilyExportedIdentifiers
func (v multiline) AfterApply() error { prettyConfig.Multiline = bool(v); return nil }

// AfterApply ...
//goland:noinspection GoUnnecessarilyExportedIdentifiers
func (v tty) AfterApply() error { prettyConfig.TTY = bool(v); return nil }

// AfterApply ...
//goland:noinspection GoUnnecessarilyExportedIdentifiers
func (v width) AfterApply() error { prettyConfig.Width = uint(v); return nil }

// AfterApply ...
//goland:noinspection GoUnnecessarilyExportedIdentifiers
func (v wrap) AfterApply() error { prettyConfig.Wrap = bool(v); return nil }

// AfterApply ...
//goland:noinspection GoUnnecessarilyExportedIdentifiers
func (v themeName) AfterApply() error {
	thm, ok := theme.Registry[string(v)]
	if !ok {
		return errors2.Errorf("unknown theme: %s", v)
	}
	prettyConfig.Theme = thm
	return nil
}

// AfterApply ...
//goland:noinspection GoUnnecessarilyExportedIdentifiers
func (t *timeFormat) AfterApply() error {
	format := strings.ToLower(string(*t))
	switch format {
	case "c":
		*t = time.ANSIC
	case "unix":
		*t = time.UnixDate
	case "ruby":
		*t = time.RubyDate
	case "rfc822":
		*t = time.RFC822
	case "rfc822-utc":
		*t = time.RFC822Z
	case "rfc850":
		*t = time.RFC850
	case "rfc1123":
		*t = time.RFC1123
	case "rfc1123-utc":
		*t = time.RFC1123Z
	case "rfc3339", "iso8601":
		*t = time.RFC3339
	case "rfc3339-nanos", "iso8601-nanos":
		*t = time.RFC3339Nano
	case "kitchen":
		*t = time.Kitchen
	case "stamp":
		*t = time.Stamp
	case "stamp-millis":
		*t = time.StampMilli
	case "stamp-micros":
		*t = time.StampMicro
	case "stamp-nanos":
		*t = time.StampNano
	}
	return nil
}

func regexpDecoder() kong.MapperFunc {
	return func(ctx *kong.DecodeContext, target reflect.Value) error {
		var value string
		if err := ctx.Scan.PopValueInto("regex", &value); err != nil {
			return err
		}
		r, err := regexp.Compile(value)
		if err != nil {
			return err
		}
		target.Set(reflect.ValueOf(*r))
		return nil
	}
}

func timeLikeDecoder() kong.MapperFunc {
	return func(ctx *kong.DecodeContext, target reflect.Value) error {
		var value string
		if err := ctx.Scan.PopValueInto("(time|duration)", &value); err != nil {
			return err
		}
		var t time.Time
		d, err := time.ParseDuration(value)
		if err == nil {
			t = now.Add(-d)
		} else {
			t, err = dateparse.ParseAny(value)
			if err != nil {
				return errors2.Errorf("expected duration but got %q: %s", value, err)
			}
		}
		target.Set(reflect.ValueOf(t))
		return nil
	}
}