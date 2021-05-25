package column

import (
	"couture/internal/pkg/model"
	"couture/internal/pkg/sink/pretty/config"
	"couture/internal/pkg/sink/pretty/theme"
	"couture/internal/pkg/source"
	"couture/internal/pkg/tty"
	"github.com/i582/cfmt/cmd/cfmt"
)

type sourceColumn struct {
	baseColumn
}

func newSourceColumn() sourceColumn {
	const weight = 40
	return sourceColumn{baseColumn{
		columnName:  "source",
		weightType:  weighted,
		widthWeight: weight,
	}}
}

// RegisterStyles ...
func (col sourceColumn) RegisterStyles(_ theme.Theme) {}

// RegisterSource ...
func RegisterSource(theme theme.Theme, src source.Source) {
	bgColor := <-theme.SourceColors
	fgColor := tty.Contrast(bgColor)
	cfmt.RegisterStyle(src.ID(), func(s string) string {
		return cfmt.Sprintf("{{"+string(src.Sigil())+" %s }}::"+fgColor+"|bg"+bgColor, s)
	})
}

// Format ...
func (col sourceColumn) Format(width uint, src source.Source, _ model.Event) string {
	return formatStyleOfWidth(src.ID(), width)
}

// Render ...
func (col sourceColumn) Render(_ config.Config, src source.Source, _ model.Event) []interface{} {
	return []interface{}{src.URL().ShortForm()}
}