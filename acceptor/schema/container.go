package schema

import "github.com/sqjian/bingo/pkg/errors"

type Container interface {
	Init() errors.Error
	FInit() errors.Error
	Add([]Server) errors.Error
	Remove(string) errors.Error
	Get(...string) ([]Server, errors.Error)
}
