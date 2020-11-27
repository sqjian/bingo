package golib

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/sqjian/bingo/acceptor/schema"
	"github.com/sqjian/bingo/pkg/errors"
	"github.com/sqjian/bingo/pkg/log"
	"plugin"
	"strings"
	"sync"
)

func NewLoader(viper *viper.Viper) schema.Loader {
	return &Loader{viper: viper}
}

type Loader struct {
	sync.RWMutex
	viper   *viper.Viper
	servers []schema.Server
}

func (l *Loader) reWriteServerName(serverName string) string {
	if strings.HasSuffix(serverName, ".so") {
		return serverName
	}
	return serverName + ".so"
}
func (l *Loader) Load(serverNames ...string) ([]schema.Server, errors.Error) {
	if len(serverNames) == 0 || func() bool {
		for _, serverName := range serverNames {
			if len(serverName) == 0 {
				return true
			}
		}
		return false
	}() {
		return nil, errors.NewWithErr(fmt.Errorf("can't load server, please specify serverNames"))
	}

	l.Lock()
	defer l.Unlock()

	var servers []schema.Server
	for _, serverName := range serverNames {
		server, serverErr := l.verify(l.reWriteServerName(serverName))
		if serverErr != nil {
			return nil, serverErr
		}
		servers = append(servers, server)
	}
	return servers, nil
}

func (l *Loader) Init() errors.Error {
	return nil
}

func (l *Loader) FInit() errors.Error {
	l.Lock()
	defer l.Unlock()

	for _, server := range l.servers {
		if err := server.FInit(); err != nil {
			return err
		}
	}

	return nil
}

func (l *Loader) verify(serverName string) (schema.Server, errors.Error) {
	log.GetDefLogger().Infof("verifying server: %v\n", serverName)

	pg, err := plugin.Open(serverName)
	if err != nil {
		log.GetDefLogger().Infof("can not open server: %v,err:%v\n", serverName, err)
		return nil, errors.NewWithErr(fmt.Errorf("can not open server: %v,err:%v", serverName, err))
	}

	f, err := pg.Lookup("NewServer")
	if err != nil {
		log.GetDefLogger().Infof("Lookup NewServer: %v failed\n", serverName)
		return nil, errors.NewWithErr(err)
	}
	inst, instErr := func() (schema.Server, errors.Error) {
		NewServer, NewServerOk := f.(schema.NewServer)
		if !NewServerOk {
			log.GetDefLogger().Infof("%v->convert %T to %T failed\n", serverName, f, schema.NewServerObj)
			return nil, errors.ServerMethodNotFound
		}
		server, serverErr := NewServer(l.viper)
		if serverErr != nil {
			log.GetDefLogger().Infof("%v->NewServer failed:%v\n", serverName, serverErr)
			return nil, serverErr
		}
		initErr := server.Init()
		if initErr != nil {
			log.GetDefLogger().Infof("%v->Init failed:%v\n", serverName, initErr)
			return nil, initErr
		}
		return server, nil
	}()
	return inst, instErr
}
