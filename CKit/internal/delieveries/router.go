package delieveries

import (
	"github.com/labstack/echo/v4"
)

type Router struct {
	echo *echo.Echo
}

func NewRouter(e *echo.Echo, handler *ShipmentHandler) *Router {
	e.POST("/shipments", handler.CreateShipment)
	e.GET("/shipments/:id", handler.GetShipment)
	e.GET("/shipments", handler.ListShipments)
	e.PUT("/shipments/:id", handler.UpdateShipment)
	e.DELETE("/shipments/:id", handler.DeleteShipment)
	return &Router{echo: e}
}

func (r *Router) Start(addr string) error {
	return r.echo.Start(addr)
}
