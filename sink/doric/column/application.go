package column

import (
	"github.com/pandich/couture/model"
	"github.com/pandich/couture/schema"
	"github.com/pandich/couture/sink/layout"
	"github.com/pandich/couture/theme"
)

func newApplicationColumn(style theme.Style, layout layout.ColumnLayout) column {
	return newWeightedColumn(
		schema.Application,
		layout,
		style,
		func(event model.SinkEvent) []interface{} { return []interface{}{string(event.Application)} },
	)
}