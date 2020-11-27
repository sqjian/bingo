package entrance

import (
	"context"
	"fmt"
	"github.com/spf13/viper"
	"github.com/sqjian/bingo/acceptor/container"
	"github.com/sqjian/bingo/acceptor/loader/golib"
	"github.com/sqjian/bingo/acceptor/loader/native"
	"github.com/sqjian/bingo/pkg/log"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func newEntrance(viperInst *viper.Viper, opts ...Option) *entrance {
	e := &entrance{}
	e.viperInst = viperInst
	for _, opt := range opts {
		opt.apply(e)
	}

	if nil == e.logger {
		e.logger = log.GetDefLogger()
	}

	return e
}

type entrance struct {
	viperInst *viper.Viper
	logger    log.Logger
}

func Serve(viperInst *viper.Viper, opts ...Option) error {
	entry := newEntrance(viperInst, opts...)
	return entry.serve()
}
func (e *entrance) setupCloseHandler(cancel func()) {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		e.logger.Infof("Ctrl+C pressed in Terminal\n")
		cancel()
	}()
}
func (e *entrance) serve() error {

	e.logger.Infof("newContainer...")
	containerInst := container.NewContainer(e.viperInst)
	e.logger.Infof("containerInst init...")
	if err := containerInst.Init(); err != nil {
		e.logger.Errorf("containerInst init failed,err:%w", err)
		return err
	}
	defer containerInst.FInit()

	var allServerNames []string

	//init acceptor
	{
		// init golib acceptor
		{
			e.logger.Infof("golib.NewLoader...")
			loader := golib.NewLoader(e.viperInst)
			if err := loader.Init(); err != nil {
				e.logger.Errorf("golib.loader.Init failed,err:%w", err)
				return err
			}

			golibAcceptors := e.viperInst.GetStringSlice("acceptor.golib")
			e.logger.Infof("about to deal golibAcceptors:%v", golibAcceptors)
			if len(golibAcceptors) == 0 {
				e.logger.Infof("empty golibAcceptors.")
			} else {
				e.logger.Infof("about to load %v", golibAcceptors)
				servers, serversErr := loader.Load(golibAcceptors...)
				if serversErr != nil {
					e.logger.Errorf("golib.loader.Load failed,serversErr:%w", serversErr)
					return serversErr
				}

				addErr := containerInst.Add(servers)
				if addErr != nil {
					e.logger.Errorf("golib.loader.Add failed,addErr:%w", addErr)
					return addErr
				}
				allServerNames = append(allServerNames, golibAcceptors...)
			}
		}
		// init native acceptor
		{
			e.logger.Infof("native.NewLoader...")
			loader := native.NewLoader(e.viperInst)
			if err := loader.Init(); err != nil {
				e.logger.Errorf("native.loader.Init failed,err:%w", err)
				return err
			}
			nativeAcceptors := e.viperInst.GetStringSlice("acceptor.native")
			e.logger.Infof("about to deal nativeAcceptors:%v", nativeAcceptors)
			if len(nativeAcceptors) == 0 {
				e.logger.Infof("empty nativeAcceptors.")
			} else {
				e.logger.Infof("about to load %v", nativeAcceptors)
				servers, serversErr := loader.Load(nativeAcceptors...)
				if serversErr != nil {
					e.logger.Errorf("native.loader.Load failed,serversErr:%w", serversErr)
					return serversErr
				}

				addErr := containerInst.Add(servers)
				if addErr != nil {
					e.logger.Errorf("native.loader.Add failed,addErr:%w", addErr)
					return addErr
				}
				allServerNames = append(allServerNames, nativeAcceptors...)
			}
		}
	}

	// start server
	{
		if len(allServerNames) == 0 {
			e.logger.Infof("empty servers,exiting...")
			return nil
		}

		e.logger.Infof("about to starting servers:%v", allServerNames)
		servers, serversErr := containerInst.Get(allServerNames...)
		if serversErr != nil {
			e.logger.Errorf("containerInst.Get failed,serversErr:%w", serversErr)
			return serversErr
		}

		var wg sync.WaitGroup

		type RunStat struct {
			name string
			err  error
		}

		var chStats = make(chan RunStat, len(servers))

		var ctx, cancel = context.WithCancel(context.Background())
		e.setupCloseHandler(cancel)

		for ix, server := range servers {
			e.logger.Infof("about to starting->server_%v:%v\n", ix, server.Name())
			wg.Add(1)
			server := server
			go func(chErr chan<- RunStat) {
				defer wg.Done()
				err := server.Run(ctx)
				chErr <- RunStat{
					name: server.Name(),
					err:  err,
				}
				return
			}(chStats)
		}

		e.logger.Errorf("waiting for all servers complete")
		wg.Wait()

		e.logger.Errorf("all servers complete,chStatsLen:%v", len(chStats))

		var runStats []RunStat
	end:
		for {
			select {
			case runStat := <-chStats:
				{
					runStats = append(runStats, runStat)
					if len(chStats) == 0 {
						break end
					}
				}
			}
		}

		var runErrs []error
		for _, runStat := range runStats {
			e.logger.Errorf("status->server_%v,err:%v\n", runStat.name, runStat.err)
			if runStat.err != nil {
				runErrs = append(runErrs, runStat.err)
			}
		}
		if len(runErrs) != 0 {
			return fmt.Errorf("runErrs:%v", runErrs)
		}
	}

	return nil
}
