package cli

import (
	"fmt"
	"github.com/alecthomas/kong"
	"github.com/pkg/errors"
	"reflect"
	"regexp"
)

var (
	//coreMappers are kong.Mapper mappers exposed via kong.Option structs.
	coreMappers = []kong.Option{
		//regexp
		kong.TypeMapper(reflect.PtrTo(reflect.TypeOf(regexp.Regexp{})), regexpMapper{}),
		kong.TypeMapper(reflect.SliceOf(reflect.PtrTo(reflect.TypeOf(regexp.Regexp{}))), regexpMapper{}),
	}
)

type (
	//creator converts a string into a resource (e.g. source or sink).
	creator func(config string) interface{}
	//creators maps reflect.Type to creator.
	creators map[reflect.Type]creator
	//creatorMapper implements the kong.Mapper interface.
	creatorMapper struct {
		creators creators
	}
	//regexpMapper uses regexp.Compile to compile the specified pattern.
	regexpMapper struct{}
)

//mapper creates a new kong.Option registering a kong.Mapper for a creator for required, optional, and slice types.
func mapper(i interface{}, creator creator) []kong.Option {
	t := reflect.TypeOf(i)
	return []kong.Option{
		kong.TypeMapper(t, creatorMapper{creators: creators{reflect.PtrTo(t): creator}}),
		kong.TypeMapper(reflect.PtrTo(t), creatorMapper{creators: creators{reflect.PtrTo(t): creator}}),
		kong.TypeMapper(reflect.SliceOf(t), creatorMapper{creators: creators{reflect.SliceOf(t): creator}}),
	}
}

func (m creatorMapper) Decode(ctx *kong.DecodeContext, target reflect.Value) error {
	if ctx.Scan.Peek().Type == kong.FlagValueToken {
		token := ctx.Scan.Pop()
		switch config := token.Value.(type) {
		case string:
			creator, ok := m.creators[target.Type()]
			if !ok {
				return errors.Errorf("unknown type (%T) with config %s", token.Value, config)
			}
			value := reflect.ValueOf(creator(config))
			switch target.Kind() {
			case reflect.Slice:
				target.Set(reflect.Append(target, value))
			case reflect.Ptr:
				target.Elem().Set(value)
			default:
				target.Set(value)
			}

		default:
			return errors.Errorf("expected string but got %q (%T)", token.Value, token.Value)
		}
	}
	return nil
}

func (m regexpMapper) Decode(ctx *kong.DecodeContext, target reflect.Value) error {
	token := ctx.Scan.Pop()
	switch pattern := token.Value.(type) {
	case string:
		filter, err := regexp.Compile(pattern)
		if err != nil {
			return err
		}
		target.Set(reflect.Append(target, reflect.ValueOf(filter)))
	default:
		return fmt.Errorf("bad type %T %v", token, token)
	}
	return nil
}
