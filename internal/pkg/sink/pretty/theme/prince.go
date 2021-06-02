package theme

import (
	"couture/internal/pkg/model/level"
	"github.com/charmbracelet/lipgloss"
)

var prince = "prince"

// init prince theme
//nolint:funlen
func init() {
	register(prince, Theme{
		Timestamp:       style("#b2a73e", "#0c0a05"),
		Application:     style("#587585", "#060708"),
		Thread:          style("#594768", "#060507"),
		Class:           style("#9b99bf", "#202020"),
		MethodDelimiter: style("#493858", "#202020"),
		Method:          style("#493858", "#202020"),
		LineDelimiter:   style("#916d90", "#202020"),
		Line:            style("#bf99bd", "#202020"),
		Level: map[level.Level]lipgloss.Style{
			level.Trace: style("#000000", "#868686"),
			level.Debug: style("#000000", "#f6f6f6"),
			level.Info:  style("#000000", "#b1e200"),
			level.Warn:  style("#000000", "#ffca00"),
			level.Error: style("#000000", "#e66a00"),
		},
		Message: map[level.Level]lipgloss.Style{
			level.Trace: style("#fbfafc", "#121212"),
			level.Debug: style("#fbfafc", "#1b1b1b"),
			level.Info:  style("#fbfafc", "#181a0a"),
			level.Warn:  style("#fbfafc", "#1e180a"),
			level.Error: style("#e66a00", "#1d1005"),
		},
		Source: []lipgloss.Style{
			style("#000000", "#d5bfc2"),
			style("#000000", "#c19089"),
			style("#000000", "#debcb7"),
			style("#000000", "#a3c4d7"),
			style("#000000", "#929e83"),
			style("#000000", "#e3998b"),
			style("#000000", "#ddc7b8"),
			style("#000000", "#9f7f82"),
			style("#000000", "#e48285"),
			style("#ffffff", "#433c49"),
			style("#000000", "#82b0d6"),
			style("#000000", "#78aad6"),
			style("#000000", "#7c80d7"),
			style("#000000", "#7479e6"),
			style("#000000", "#a299c7"),
			style("#000000", "#beb5c7"),
			style("#000000", "#6998ba"),
			style("#000000", "#8470bb"),
			style("#000000", "#797bd7"),
			style("#000000", "#ce7194"),
			style("#000000", "#c46072"),
			style("#000000", "#d17ec8"),
			style("#000000", "#b47370"),
			style("#000000", "#e7826e"),
			style("#000000", "#d8847d"),
			style("#000000", "#966f69"),
			style("#000000", "#aba8c1"),
			style("#000000", "#e1d5ab"),
			style("#000000", "#71baa9"),
			style("#000000", "#d6658b"),
			style("#000000", "#dfa1c5"),
			style("#000000", "#8b80c7"),
			style("#000000", "#cd659d"),
			style("#000000", "#6ea0cb"),
			style("#000000", "#749ac7"),
			style("#000000", "#8b7c86"),
			style("#000000", "#9a94c4"),
			style("#000000", "#cc908a"),
			style("#000000", "#a1b8d9"),
			style("#000000", "#dfabcc"),
			style("#000000", "#e2be89"),
			style("#000000", "#ba6360"),
			style("#000000", "#6976be"),
			style("#000000", "#d5b5b6"),
			style("#ffffff", "#7f627b"),
			style("#000000", "#c8d672"),
			style("#000000", "#c76b78"),
			style("#000000", "#87ac8e"),
			style("#000000", "#83a394"),
			style("#000000", "#bb73c0"),
			style("#ffffff", "#9d645d"),
			style("#000000", "#9e958e"),
			style("#000000", "#d2afa0"),
			style("#000000", "#e1bd8e"),
			style("#000000", "#a8c4a6"),
			style("#000000", "#998e99"),
			style("#000000", "#74a380"),
			style("#000000", "#6f8ec5"),
			style("#000000", "#d9c7a0"),
			style("#000000", "#e263c7"),
			style("#000000", "#b7c66f"),
			style("#000000", "#ba6d76"),
			style("#000000", "#7b7bc4"),
			style("#ffffff", "#ae5483"),
			style("#000000", "#6fa69d"),
			style("#000000", "#e3d590"),
			style("#000000", "#b893c7"),
			style("#000000", "#dda6ca"),
			style("#000000", "#e1d5ae"),
			style("#000000", "#9fbfce"),
			style("#000000", "#bdc781"),
			style("#000000", "#d58383"),
			style("#000000", "#e2ac98"),
			style("#000000", "#c9626b"),
			style("#000000", "#dd75b0"),
			style("#000000", "#aecdca"),
			style("#000000", "#c7626a"),
			style("#000000", "#dcb37e"),
			style("#000000", "#9a94ad"),
			style("#000000", "#da825d"),
			style("#000000", "#c25c75"),
			style("#000000", "#d795b2"),
			style("#000000", "#7ba5c7"),
			style("#000000", "#76b1bf"),
			style("#000000", "#ada4d6"),
			style("#000000", "#82926b"),
			style("#000000", "#cdc783"),
			style("#ffffff", "#8b5c99"),
			style("#000000", "#d35e5f"),
			style("#000000", "#ac6394"),
			style("#000000", "#e2d27c"),
			style("#000000", "#dcac7d"),
			style("#000000", "#89a297"),
			style("#000000", "#c980be"),
			style("#000000", "#9d69b8"),
			style("#000000", "#dfadb5"),
			style("#000000", "#97bdd6"),
			style("#000000", "#9ec0c7"),
			style("#000000", "#9a81bf"),
			style("#000000", "#789280"),
			style("#000000", "#b0b088"),
			style("#000000", "#c98cb9"),
			style("#000000", "#d39292"),
			style("#000000", "#ac6f7d"),
			style("#000000", "#e5c970"),
			style("#000000", "#d6a494"),
			style("#ffffff", "#5c65a0"),
			style("#000000", "#6eab9c"),
			style("#000000", "#c070a9"),
			style("#000000", "#6a79c7"),
			style("#000000", "#e69971"),
			style("#000000", "#b6a985"),
			style("#000000", "#e68f6f"),
			style("#000000", "#e77478"),
			style("#000000", "#debc8e"),
			style("#000000", "#cd9bca"),
			style("#ffffff", "#5e5d6c"),
			style("#000000", "#e77a6e"),
			style("#000000", "#6c9cc2"),
			style("#000000", "#e0b6b0"),
			style("#000000", "#bbb8d6"),
			style("#000000", "#d76265"),
			style("#000000", "#ddbed5"),
			style("#000000", "#5d878e"),
			style("#000000", "#de87e4"),
			style("#000000", "#dc9fa2"),
			style("#ffffff", "#9d5d94"),
			style("#ffffff", "#875ac1"),
			style("#000000", "#c6a6d1"),
			style("#000000", "#e263c7"),
			style("#000000", "#e66984"),
			style("#000000", "#c98877"),
			style("#000000", "#d4ad87"),
			style("#ffffff", "#846d68"),
			style("#000000", "#d76acb"),
			style("#000000", "#d65d82"),
			style("#000000", "#de5d70"),
			style("#000000", "#e87e64"),
			style("#000000", "#b9629c"),
			style("#000000", "#73b9c7"),
			style("#ffffff", "#8763ad"),
			style("#000000", "#e195af"),
			style("#000000", "#e8606b"),
			style("#000000", "#93d78f"),
			style("#000000", "#a1c3bb"),
			style("#000000", "#a97371"),
			style("#000000", "#96807b"),
			style("#000000", "#78b9a9"),
			style("#000000", "#e084e5"),
			style("#000000", "#bfb0bd"),
			style("#000000", "#93bfd9"),
			style("#000000", "#d4cabf"),
			style("#000000", "#e5b776"),
			style("#000000", "#e77172"),
			style("#000000", "#cb998e"),
			style("#000000", "#608b99"),
			style("#000000", "#d3a7cd"),
			style("#000000", "#df8caf"),
			style("#000000", "#c8bdc9"),
			style("#000000", "#5c8584"),
			style("#000000", "#cda19a"),
			style("#000000", "#8fc1d7"),
			style("#ffffff", "#7755c0"),
			style("#000000", "#e3d590"),
			style("#ffffff", "#856361"),
			style("#000000", "#c77270"),
			style("#ffffff", "#8c598c"),
			style("#000000", "#9572ae"),
			style("#000000", "#8e7fc5"),
			style("#000000", "#df6f9f"),
			style("#000000", "#e2999a"),
			style("#ffffff", "#945fa3"),
			style("#000000", "#ded5e5"),
			style("#000000", "#908fbb"),
			style("#000000", "#e36aa8"),
			style("#000000", "#e37894"),
			style("#000000", "#be9dd0"),
			style("#000000", "#e0c894"),
			style("#000000", "#bfc498"),
			style("#000000", "#e5a579"),
		},
	})
}
