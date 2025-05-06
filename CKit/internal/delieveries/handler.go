package delieveries

import (
	"CKit/internal/entity"
	"CKit/internal/middleware"
	"CKit/internal/services/shipment"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

// ShipmentHandler aggregates shipment use cases.
type ShipmentHandler struct {
	CreateUC *shipment.CreateShipmentUseCase
	GetUC    *shipment.GetShipmentUseCase
	ListUC   *shipment.ListShipmentsUseCase
	UpdateUC *shipment.UpdateShipmentUseCase
	DeleteUC *shipment.DeleteShipmentUseCase
}

// NewShipmentHandler constructs a handler with all use cases.
func NewShipmentHandler(
	create *shipment.CreateShipmentUseCase,
	get *shipment.GetShipmentUseCase,
	list *shipment.ListShipmentsUseCase,
	update *shipment.UpdateShipmentUseCase,
	deleteUC *shipment.DeleteShipmentUseCase,
) *ShipmentHandler {
	return &ShipmentHandler{
		CreateUC: create,
		GetUC:    get,
		ListUC:   list,
		UpdateUC: update,
		DeleteUC: deleteUC,
	}
}

// CreateShipment handles POST /shipments
func (h *ShipmentHandler) CreateShipment(c echo.Context) error {
	authInfo := c.Get("AuthInfo").(middleware.AuthInfo)

	if authInfo.Role == "driver" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid role"})
	}

	var req entity.CreateShipmentRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid payload"})
	}
	ctx := c.Request().Context()
	subInfo := c.Get("SubscriptionInfo").(middleware.SubscriptionInfo)
	sub, err := h.CreateUC.Execute(ctx, authInfo.UserId, &subInfo, &req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusCreated, sub)
}

// GetShipment handles GET /shipments/:id
func (h *ShipmentHandler) GetShipment(c echo.Context) error {
	id := c.Param("id")
	ctx := c.Request().Context()
	s, err := h.GetUC.Execute(ctx, id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "not found"})
	}
	return c.JSON(http.StatusOK, s)
}

// ListShipments handles GET /shipments
func (h *ShipmentHandler) ListShipments(c echo.Context) error {
	authInfo := c.Get("AuthInfo").(middleware.AuthInfo)
	ctx := c.Request().Context()
	var list []*entity.Shipment
	var err error

	if authInfo.Role == "customer" {
		list, err = h.ListUC.ExecuteCustomer(ctx, authInfo.UserId)
	} else {
		list, err = h.ListUC.ExecuteDriver(ctx, authInfo.UserId)
	}
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, list)
}

// UpdateShipment handles PUT /shipments/:id
func (h *ShipmentHandler) UpdateShipment(c echo.Context) error {
	authInfo := c.Get("AuthInfo").(middleware.AuthInfo)

	if authInfo.Role == "driver" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid role"})
	}
	id := c.Param("id")
	var req entity.Shipment
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid payload"})
	}
	req.ID = id
	ctx := c.Request().Context()
	updated, err := h.UpdateUC.Execute(ctx, &req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, updated)
}

// DeleteShipment handles DELETE /shipments/:id
func (h *ShipmentHandler) DeleteShipment(c echo.Context) error {
	authInfo := c.Get("AuthInfo").(middleware.AuthInfo)

	if authInfo.Role == "driver" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid role"})
	}
	id := c.Param("id")
	ctx := c.Request().Context()
	if err := h.DeleteUC.Execute(ctx, id); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.NoContent(http.StatusNoContent)
}

func (h *ShipmentHandler) updateShipmentStatus(c echo.Context, status entity.ShipmentStatus) error {
	authInfo := c.Get("AuthInfo").(middleware.AuthInfo)

	if authInfo.Role == "customer" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid role"})
	}
	id := c.Param("id")
	ctx := c.Request().Context()

	shipment, err := h.GetUC.Execute(ctx, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	fmt.Printf("shipment: %s", shipment)

	shipment.Status = status
	shipment.PickerID = authInfo.UserId

	updated, err := h.UpdateUC.Execute(ctx, shipment)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, updated)
}

// TakeShipment handles PUT /shipments/take/:id
func (h *ShipmentHandler) TakeShipment(c echo.Context) error {
	return h.updateShipmentStatus(c, entity.StatusInTransit)
}

// CancelShipment handles PUT /shipments/cancel/:id
func (h *ShipmentHandler) CancelShipment(c echo.Context) error {
	return h.updateShipmentStatus(c, entity.StatusCancelled)
}

// DoneShipment handles PUT /shipments/done/:id
func (h *ShipmentHandler) DoneShipment(c echo.Context) error {
	return h.updateShipmentStatus(c, entity.StatusDelivered)
}
