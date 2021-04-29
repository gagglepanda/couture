package cli

import (
	"couture/internal/pkg/model"
	"github.com/alecthomas/kong"
	"log"
	"time"
)

func init() {
	coreCli.Plugins = kong.Plugins{&sinkCLI, &sourceCLI}
}

var (
	//coreOptions contain the core kong cli options.
	coreOptions = []kong.Option{
		kong.Name("couture"),
		kong.Description("Tail multiple log sources."),
		kong.ShortUsageOnError(),
		kong.ConfigureHelp(kong.HelpOptions{
			Tree:     true,
			Indenter: kong.TreeIndenter,
		}),
	}

	//coreCli contains all global arguments.
	coreCli struct {
		AwsRegion  string `group:"AWS Options" help:"AWS region" default:"us-west-2" name:"aws-region" aliases:"region" env:"AWS_REGION"`
		AwsProfile string `group:"AWS Options" help:"AWS profile" default:"integration" name:"aws-profile" aliases:"profile" env:"AWS_PROFILE"`

		Quiet       bool `group:"Display Options" help:"Only log lines are displayed. Header and diagnostics are suppressed." name:"quiet" aliases:"silent" short:"q" xor:"verbosity"`
		Verbose     bool `group:"Display Options" help:"Display additional diagnostic data." name:"verbose" short:"v" xor:"verbosity"`
		ClearScreen bool `group:"Display Options" help:"Clear screen prior to start." name:"clear" default:"true" negatable:"true"`
		ShowPrefix  bool `group:"Display Options" help:"Display a prefix before each log line indicting its source." name:"prefix" default:"true" negatable:"true"`
		ShortNames  bool `group:"Display Options" help:"Display a abbreviated source names." name:"short-names" aliases:"short" default:"true" negatable:"true"`

		Follow      bool          `group:"Behavioral Options" help:"Follow the logs." default:"true" name:"follow" short:"f" negatable:"true"`
		PollCadence time.Duration `group:"Behavioral Options" help:"How long to sleep between polls. (Applies only to some sources.)" default:"2s" name:"interval" aliases:"sleep" short:"i"`
		LineCount   uint32        `group:"Behavioral Options" help:"How many lines of history to include. (Applies only to some sources.)" default:"20" name:"lines" aliases:"count" short:"c"`
		Since       time.Duration `group:"Behavioral Options" help:"How far back to search for events." default:"5m" name:"since" aliases:"back,lookback" short:"b"`

		Patterns []string    `group:"Filtering Options" help:"Filter patterns." name:"filter" short:"f" sep:","`
		LogLevel model.Level `group:"Filtering Options" help:"Minimum log level to display (${enum})." default:"DEBUG" name:"log-level" aliases:"level" short:"l" enum:"ERROR,WARN,INFO,DEBUG,TRACE"`

		kong.Plugins
	}
)

//Parse parses all arguments, including all sources and sinks.
func Parse() *kong.Context {
	var opts []kong.Option
	opts = append(opts, coreOptions...)
	opts = append(opts, sourceMappers...)
	opts = append(opts, sinkMappers...)
	ctx := kong.Parse(&coreCli, opts...)
	log.Printf("%+v\n", coreCli)
	return ctx
}