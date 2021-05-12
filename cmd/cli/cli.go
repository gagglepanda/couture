package cli

import (
	"couture/internal/pkg/sink/json"
	"couture/internal/pkg/sink/pretty"
	"couture/internal/pkg/sink/simple"
	"couture/pkg/model"
	"github.com/alecthomas/kong"
	"net/url"
	"os"
	"regexp"
	"strings"
)

// MustParseOptions parses sources arguments, including sources registry and sinks.
func MustParseOptions() ([]interface{}, error) {
	var exampleURLs []string
	for _, src := range sources {
		exampleURLs = append(exampleURLs, src.ExampleURLs...)
	}
	var kongOptions []kong.Option
	kongOptions = append(kongOptions,
		kong.Name(os.Args[0]),
		kong.Description(strings.Join([]string{
			"Overview:\n",
			"Tails one or more event sources. When providing a CloudFormation stack, resources " +
				"are recursively analyzed until all loggable entities are found. " +
				"This includes the stack events of the stack itself, as well as any log groups " +
				"its entities contain.\n",
			"Supported URL Formats:\n",
			"\t" + strings.Join(exampleURLs, "\n\t"),
		}, "\n")),
		kong.ConfigureHelp(kong.HelpOptions{
			Tree:      true,
			Indenter:  kong.TreeIndenter,
			FlagsLast: true,
			Compact:   true,
		}),
		kong.ShortUsageOnError(),
	)
	kongOptions = append(kongOptions, sinkMappers...)
	kongOptions = append(kongOptions, coreMappers...)
	parser, err := kong.New(&cli, kongOptions...)
	if err != nil {
		return nil, err
	}
	_, err = parser.Parse(os.Args[1:])
	if err != nil {
		return nil, err
	}

	var options []interface{}
	options = append(options, *configuredSink())
	options = append(options, filterOptions()...)
	options = append(options, displayOptions()...)
	sources, err := configuredSources()
	if err != nil {
		return nil, err
	}
	options = append(options, sources...)
	return options, nil
}

// cli is the container for sources CLI argument plugins.
//nolint:lll
var cli struct {
	Verbosity uint `group:"Display" help:"Display additional diagnostic data." name:"verbose" short:"v" xor:"verbosity" type:"counter" placeholder:"<level>" enum:"0,1,2,3,4,5,6" env:"COUTURE_VERBOSITY"`
	Quiet     bool `group:"Display" help:"Display no diagnostic data." name:"quiet" short:"q" xor:"verbosity"`
	Wrap      uint `group:"Display" help:"Wrap output to the specified width." name:"wrap" short:"w" placeholder:"<width>" env:"COUTURE_WRAP"`
	Emphasis  bool `group:"Display" help:"Emphasize message intensity based upon log level." name:"emphasis" short:"e" default:"true" negatable:"true"`
	Inactive  bool `group:"Display" name:"inactive" help:"Include notifications for every log group, even ones without new events." negatable:"true" default:"false"`
	Billing   bool `group:"Display" name:"billing" help:"Exclude billing report data." negatable:"true" default:"true"`
	Raw       bool `group:"Display" name:"raw" help:"Exclude sources non-JSON events." negatable:"true" default:"true"`

	IncludeFilters []*regexp.Regexp `group:"Filter" help:"Include filter regular expressions. Always performed before excludes." name:"include" placeholder:"<regex>" short:"i" sep:"none"`
	ExcludeFilters []*regexp.Regexp `group:"Filter" help:"Exclude filter regular expressions. Always performed after includes." name:"exclude" placeholder:"<regex>" short:"x" sep:"none"`

	Level  model.Level  `group:"Filter" help:"Minimum log level to display (${enum})." default:"trace" name:"level" short:"l" enum:"error,warn,info,debug,trace" env:"COUTURE_LOG_LEVEL"`
	Simple *simple.Sink `group:"Sink" help:"Dump string representation of event." name:"simple" short:"s" placeholder:"full" xor:"sink"`
	JSON   *json.Sink   `group:"Sink" help:"Dump Sink representation of event." name:"json" short:"j" placeholder:"pretty" xor:"sink"`
	Pretty *pretty.Sink `group:"Sink" help:"Pretty columnar output. To explicitly disable color set the NO_COLOR environment variable." name:"pretty" short:"p" xor:"sink"`

	Sources []url.URL `arg:"true" group:"Source" name:"source_url" help:"One ore more log sources. (See: Supported URL Formats)"`

	cliValidator
}
