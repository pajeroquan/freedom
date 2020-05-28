package general

import (
	"reflect"
	"time"

	stdContext "context"

	iris "github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"github.com/kataras/iris/v12/core/memstore"
)

func newWorkerHandle() context.Handler {
	return func(ctx context.Context) {
		work := newWorker(ctx)
		ctx.Values().Set(WorkerKey, work)
		ctx.Next()
		if !work.IsRecycle() {
			return
		}
		work.logger = nil
		work.ctx = nil
		ctx.Values().Reset()
	}
}

func newWorker(ctx iris.Context) *worker {
	work := new(worker)
	work.freeServices = make([]interface{}, 0)
	work.coms = make(map[reflect.Type]interface{})
	work.ctx = ctx
	work.bus = newBus(ctx.Request().Header)
	work.stdCtx = ctx.Request().Context()
	work.time = time.Now()
	work.recycle = true
	HandleBusMiddleware(work)
	return work
}

// worker .
type worker struct {
	ctx          iris.Context
	freeServices []interface{}
	coms         map[reflect.Type]interface{}
	logger       Logger
	bus          *Bus
	stdCtx       stdContext.Context
	time         time.Time
	values       memstore.Store
	recycle      bool
}

// Ctx .
func (rt *worker) IrisContext() iris.Context {
	return rt.ctx
}

// Context .
func (rt *worker) Context() stdContext.Context {
	return rt.stdCtx
}

// WithContext .
func (rt *worker) WithContext(ctx stdContext.Context) {
	rt.stdCtx = ctx
}

// StartTime .
func (rt *worker) StartTime() time.Time {
	return rt.time
}

// Logger .
func (rt *worker) Logger() Logger {
	if rt.logger == nil {
		l := rt.values.Get("logger_impl")
		if l == nil {
			rt.logger = rt.ctx.Application().Logger()
		} else {
			rt.logger = l.(Logger)
		}
	}
	return rt.logger
}

// Store .
func (rt *worker) Store() *memstore.Store {
	return &rt.values
}

// Bus .
func (rt *worker) Bus() *Bus {
	return rt.bus
}

// CloseRecycle .
func (rt *worker) CloseRecycle() {
	rt.recycle = false
}

// SetRecycle .
func (rt *worker) IsRecycle() bool {
	return rt.recycle
}