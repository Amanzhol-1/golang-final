package delieveries

import (
	"CKit/internal/middleware"

	"github.com/labstack/echo/v4"
)

type Router struct {
	echo *echo.Echo
}

func NewRouter(e *echo.Echo, handler *ShipmentHandler, authChecker middleware.AuthChecker, subChecker middleware.SubscriptionChecker) *Router {
	grp := e.Group("/shipments")
	grp.Use(middleware.RequireAuth(authChecker))
	grp.Use(middleware.RequireSubscription(subChecker))

	grp.POST("", handler.CreateShipment)
	grp.GET("/:id", handler.GetShipment)
	grp.GET("", handler.ListShipments)
	grp.PUT("/:id", handler.UpdateShipment)
	grp.DELETE("/:id", handler.DeleteShipment)

	grp.PUT("/take/:id", handler.TakeShipment)
	grp.PUT("/cancel/:id", handler.CancelShipment)
	grp.PUT("/done/:id", handler.DoneShipment)

	return &Router{echo: e}
}

func (r *Router) Start(addr string) error {
	return r.echo.Start(addr)
}
