package doric

import (
	"github.com/i582/cfmt/cmd/cfmt"
	"github.com/pandich/couture/internal/pkg/io"
	"github.com/pandich/couture/internal/pkg/model"
	"github.com/pandich/couture/internal/pkg/sink"
	"github.com/pandich/couture/internal/pkg/sink/doric/column"
	"github.com/pandich/couture/internal/pkg/source"
)

const sinkName = "doric"

// doricSink provides render output.
type doricSink struct {
	terminalWidth uint
	table         *column.Table
	config        sink.Config
	out           chan string
}

// New provides a configured doricSink sink.
func New(config sink.Config) sink.Sink {
	return &doricSink{
		terminalWidth: config.EffectiveTerminalWidth(),
		table:         column.NewTable(config),
		config:        config,
		out:           io.NewOut(sinkName, config.Out),
	}
}

// Init ...
func (snk doricSink) Init(sources []*source.Source) {
	switch {
	case snk.config.Color != nil && !*snk.config.Color:
		cfmt.DisableColors()
	case snk.config.EffectiveIsTTY():
		cfmt.EnableColors()
	default:
		cfmt.DisableColors()
	}
	var sourceColors = map[model.SourceURL]string{}
	for _, src := range sources {
		consistentColors := *snk.config.ConsistentColors
		style := snk.config.Theme.SourceStyle(consistentColors, *src)
		sourceColors[(*src).URL()] = style.Bg
		column.RegisterSourceStyle(style, snk.config.Layout.Source, *src)
	}
}

// Accept ...
func (snk doricSink) Accept(event model.SinkEvent) error {
	snk.out <- snk.table.Render(event)
	return nil
}
