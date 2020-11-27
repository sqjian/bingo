package schema

import (
	"context"
	"github.com/sqjian/bingo/pkg/errors"
)

type Server interface {
	Name() string
	Init() errors.Error
	FInit() errors.Error
	Run(ctx context.Context) errors.Error
}
