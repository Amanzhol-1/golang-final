package delieveries

import (
	"CKit/internal/middleware"

	"github.com/labstack/echo/v4"
)

type Router struct {
	echo *echo.Echo
}

func NewRouter(e *echo.Echo, handler *ShipmentHandler, checker middleware.SubscriptionChecker) *Router {
	grp := e.Group("/shipments")
	grp.Use(middleware.RequireSubscription(checker))

	grp.POST("", handler.CreateShipment)
	grp.GET("/:id", handler.GetShipment)
	grp.GET("", handler.ListShipments)
	grp.PUT("/:id", handler.UpdateShipment)
	grp.DELETE("/:id", handler.DeleteShipment)
	return &Router{echo: e}
}

func (r *Router) Start(addr string) error {
	return r.echo.Start(addr)
}
