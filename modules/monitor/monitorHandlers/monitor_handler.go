package monitorHandlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jamee5e/jame-shop-tutorial/config"
	"github.com/jamee5e/jame-shop-tutorial/modules/entities"
	"github.com/jamee5e/jame-shop-tutorial/modules/monitor"
)

type IMonitorHandlers interface {
	HealthCheck(c *fiber.Ctx) error
}

type monitorHandlers struct {
	cfg config.IConfig
}

func MonitorHandler(cfg config.IConfig) IMonitorHandlers {
	return &monitorHandlers{
		cfg: cfg,
	}
}

func (h *monitorHandlers) HealthCheck(c *fiber.Ctx) error {
	res := &monitor.Monitor{
		Name:    h.cfg.App().Name(),
		Version: h.cfg.App().Version(),
	}
	return entities.NewResponse(c).Success(fiber.StatusOK, res).Res()

}
