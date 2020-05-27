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
		rt := newWorker(ctx)
		ctx.Values().Set(WorkerKey, rt)
		ctx.Next()
		ctx.Values().Reset()
		rt.ctx = nil
		rt.store = nil
	}
}

func newWorker(ctx iris.Context) *worker {
	rt := new(worker)
	rt.freeServices = make([]interface{}, 0)
	rt.coms = make(map[reflect.Type]interface{})
	rt.ctx = ctx
	rt.store = ctx.Values()
	rt.bus = newBus(ctx.Request().Header)
	rt.stdCtx = ctx.Request().Context()
	rt.time = time.Now()
	return rt
}

// worker .
type worker struct {
	ctx          iris.Context
	freeServices []interface{}
	coms         map[reflect.Type]interface{}
	store        *memstore.Store
	logger       Logger
	bus          *Bus
	stdCtx       stdContext.Context
	time         time.Time
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
		l := rt.ctx.Values().Get("logger_impl")
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
	return rt.ctx.Values()
}

// Bus .
func (rt *worker) Bus() *Bus {
	return rt.bus
}
