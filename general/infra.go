package general

import (
	"github.com/8treenet/freedom/infra/requests"
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
)

// Infra .
type Infra struct {
	Worker Worker `json:"-"`
}

// BeginRequest .子实现多态
func (c *Infra) BeginRequest(rt Worker) {
	c.Worker = rt
}

// DB .
func (c *Infra) DB() (db *gorm.DB) {
	return globalApp.Database.db
}

// Redis .
func (c *Infra) Redis() redis.Cmdable {
	return globalApp.Cache.client
}

// GetOther .
func (repo *Infra) GetOther(obj interface{}) {
	globalApp.other.get(obj)
	return
}

// NewHttpRequest, transferBus : Whether to pass the context, turned on by default. Typically used for tracking internal services.
func (c *Infra) NewHttpRequest(url string, transferBus ...bool) Request {
	req := requests.NewHttpRequest(url)
	if len(transferBus) > 0 && !transferBus[0] {
		return req
	}

	bus := c.Worker.Bus()
	head := bus.Header
	cloneHead := bus.Header.Clone()
	bus.Header = cloneHead
	HandleBusMiddleware(c.Worker)
	bus.Header = head
	req.SetHeader(cloneHead)
	return req
}

// NewH2CRequest, transferBus : Whether to pass the context, turned on by default. Typically used for tracking internal services.
func (c *Infra) NewH2CRequest(url string, transferBus ...bool) Request {
	req := requests.NewH2CRequest(url)
	if len(transferBus) > 0 && !transferBus[0] {
		return req
	}

	bus := c.Worker.Bus()
	head := bus.Header
	cloneHead := bus.Header.Clone()
	bus.Header = cloneHead
	HandleBusMiddleware(c.Worker)
	bus.Header = head
	req.SetHeader(cloneHead)
	return req
}
