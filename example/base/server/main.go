package main

import (
	"github.com/8treenet/freedom"
	_ "github.com/8treenet/freedom/example/base/adapter/controllers"
	"github.com/8treenet/freedom/example/base/infra/config"
	"github.com/8treenet/freedom/infra/requests"
	"github.com/8treenet/freedom/middleware"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	app := freedom.NewApplication()
	installMiddleware(app)
	addrRunner := app.CreateRunner(config.Get().App.Other["listen_addr"].(string))
	app.Run(addrRunner, *config.Get().App)
}

func installMiddleware(app freedom.Application) {
	app.InstallMiddleware(middleware.NewRecover())
	app.InstallMiddleware(middleware.NewTrace("x-request-id"))
	app.InstallMiddleware(middleware.NewRequestLogger("x-request-id", true))
	app.InstallMiddleware(middleware.NewRuntimeLogger("x-request-id"))
	app.InstallBusMiddleware(middleware.NewLimiter())
	requests.InstallPrometheus(config.Get().App.Other["service_name"].(string), freedom.Prometheus())
}
