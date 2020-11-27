package entrance

import (
	"github.com/sqjian/bingo/pkg/log"
)

type Option interface {
	apply(e *entrance)
}

type optionFunc func(e *entrance)

func (fn optionFunc) apply(e *entrance) {
	fn(e)
}

func WithAcceptorLogger(logger log.Logger) Option {
	return optionFunc(func(e *entrance) {
		e.logger = logger
	})
}
