package schema

import (
	"github.com/sqjian/bingo/pkg/errors"
)

type Container interface {
	Init() errors.Error
	FInit() errors.Error
	Add([]Plugin) errors.Error
	Remove(string) errors.Error
	Get(...string) ([]Plugin, errors.Error)
}
