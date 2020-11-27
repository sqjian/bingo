package tool

import (
	"bytes"
	"github.com/sqjian/bingo/pkg/log"
	"sync"
)

type PluginToolsImpl struct {
	BufferImpl
	LoggerImpl
	CarpoolImpl
}

type BufferImpl struct {
	sm sync.Map
}

func (t *BufferImpl) Set(k, v interface{}) {
	t.sm.Store(k, v)
}

func (t *BufferImpl) Get(k interface{}) (value interface{}, ok bool) {
	return t.sm.Load(k)
}

type LoggerImpl struct {
}

func (t *LoggerImpl) Logger() log.Logger {
	return log.GetDefLogger()
}

type CarpoolImpl struct {
	buf bytes.Buffer
}

func (t *CarpoolImpl) SendBack(buf []byte) {
	t.buf.Write(buf)
}

func (t *CarpoolImpl) GetBackData() []byte {
	return t.buf.Bytes()
}
