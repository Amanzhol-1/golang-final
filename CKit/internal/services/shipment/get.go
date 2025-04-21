package shipment

import (
	"CKit/internal/entity"
	repository "CKit/internal/repository/shipment"
	"context"
)

// GetShipmentUseCase handles fetching a shipment by ID.
type GetShipmentUseCase struct {
	repo repository.ShipmentRepository
}

// NewGetShipmentUseCase initializes the use case.
func NewGetShipmentUseCase(r repository.ShipmentRepository) *GetShipmentUseCase {
	return &GetShipmentUseCase{repo: r}
}

// Execute retrieves a shipment by its ID.
func (uc *GetShipmentUseCase) Execute(ctx context.Context, id string) (*entity.Shipment, error) {
	return uc.repo.FindByID(ctx, id)
}
