package local

import (
	"couture/internal/pkg/model"
	"couture/internal/pkg/source"
	"couture/internal/pkg/source/pipe"
	"os/exec"
	"path/filepath"
	"reflect"
	"sync"
)

// Metadata ...
func Metadata() source.Metadata {
	return source.Metadata{
		Name:      "Local File",
		Type:      reflect.TypeOf(fileSource{}),
		CanHandle: func(url model.SourceURL) bool { return url.Scheme == "file" },
		Creator: func(sourceURL model.SourceURL) (*interface{}, error) {
			var i interface{} = newSource(sourceURL)
			return &i, nil
		},
		ExampleURLs: []string{"file://<path>"},
	}
}

// fileSource ...
type fileSource struct {
	*source.Pushing
	filename string
}

// newSource ...
func newSource(sourceURL model.SourceURL) *source.Pushable {
	sourceURL.Normalize()
	var src source.Pushable = fileSource{
		Pushing:  source.New(sourceURL),
		filename: sourceURL.Path,
	}
	return &src
}

// Start ...
func (src fileSource) Start(wg *sync.WaitGroup, running func() bool, callback func(event model.Event)) error {
	// get the safe path to the file
	path, err := filepath.Abs(src.filename)
	if err != nil {
		return err
	}

	// create the command
	tail := exec.Command("tail", "-F", path)

	// capture its output
	out, err := tail.StdoutPipe()
	if err != nil {
		return err
	}

	// and start it
	if err = tail.Start(); err != nil {
		return err
	}
	return pipe.Start(wg, running, callback, func() {}, out)
}