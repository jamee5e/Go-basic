package servers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jamee5e/jame-shop-tutorial/modules/middlewares/middlewaresHandlers"
	"github.com/jamee5e/jame-shop-tutorial/modules/middlewares/middlewaresRepositories"
	"github.com/jamee5e/jame-shop-tutorial/modules/middlewares/middlewaresUsecases"
	"github.com/jamee5e/jame-shop-tutorial/modules/monitor/monitorHandlers"
)

type IModuleFactory interface {
	MonitorModule()
}

type moduleFactory struct {
	router fiber.Router
	server *server
	mid    middlewaresHandlers.IMiddlewaresHandler
}

func InitModule(r fiber.Router, s *server, mid middlewaresHandlers.IMiddlewaresHandler) IModuleFactory {
	return &moduleFactory{
		router: r,
		server: s,
		mid:    mid,
	}
}

func InitMiddlewares(s *server) middlewaresHandlers.IMiddlewaresHandler {
	repository := middlewaresRepositories.MiddlewareRepository(s.db)
	usecase := middlewaresUsecases.MiddlewaresUsercase(repository)
	return middlewaresHandlers.MiddlewaresHandler(s.cfg, usecase)
}

func (m *moduleFactory) MonitorModule() {
	handler := monitorHandlers.MonitorHandler(m.server.cfg)

	m.router.Get("/", handler.HealthCheck)
}
