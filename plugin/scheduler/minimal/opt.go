package minimal

func NewSchedulerOpts() *schedulerOpts {
	return &schedulerOpts{kvs: make(map[string]interface{})}
}

type schedulerOpts struct {
	kvs map[string]interface{}
}

type schedulerOpt func(*schedulerOpts)

func WithCarpoolOpt(kvs map[string]interface{}) schedulerOpt {
	return func(in *schedulerOpts) {
		in.kvs = kvs
	}
}
