package manager

import (
	"couture/pkg/model/level"
	"regexp"
	"time"
)

// NoWrap ...
const NoWrap = 0

// SinceOption ...
func SinceOption(t time.Time) interface{} {
	return baseOption{applier: func(options *managerOptions) {
		options.since = &t
	}}
}

// VerboseDisplayOption ...
func VerboseDisplayOption(level level.Level) interface{} {
	return baseOption{applier: func(options *managerOptions) {
		options.level = level
	}}
}

// FilterOption ...
func FilterOption(includeFilters []*regexp.Regexp, excludeFilters []*regexp.Regexp) interface{} {
	return baseOption{applier: func(options *managerOptions) {
		options.includeFilters = includeFilters
		options.excludeFilters = excludeFilters
	}}
}

// LogLevelOption ...
func LogLevelOption(level level.Level) interface{} {
	return baseOption{applier: func(options *managerOptions) {
		options.level = level
	}}
}

// WrapOption ...
func WrapOption(width uint) interface{} {
	return baseOption{applier: func(options *managerOptions) {
		if width > 0 {
			options.wrap = &width
		}
	}}
}

type (
	// managerOptions
	managerOptions struct {
		level          level.Level
		wrap           *uint
		since          *time.Time
		includeFilters []*regexp.Regexp
		excludeFilters []*regexp.Regexp
	}

	// option is an entity capable of mutating the state of a managerOptions struct.
	option interface {
		Apply(options *managerOptions)
	}

	baseOption struct {
		applier func(*managerOptions)
	}
)

// Apply ...
func (opt baseOption) Apply(mgrOptions *managerOptions) {
	opt.applier(mgrOptions)
}
