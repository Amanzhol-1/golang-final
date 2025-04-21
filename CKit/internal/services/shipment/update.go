package shipment

import (
	"CKit/internal/entity"
	repository "CKit/internal/repository/shipment"
	"context"
)

// UpdateShipmentUseCase handles updating a shipment.
type UpdateShipmentUseCase struct {
	repo repository.ShipmentRepository
}

// NewUpdateShipmentUseCase initializes the use case.
func NewUpdateShipmentUseCase(r repository.ShipmentRepository) *UpdateShipmentUseCase {
	return &UpdateShipmentUseCase{repo: r}
}

// Execute updates the given shipment.
func (uc *UpdateShipmentUseCase) Execute(ctx context.Context, shipment *entity.Shipment) (*entity.Shipment, error) {
	if err := uc.repo.Update(ctx, shipment); err != nil {
		return nil, err
	}
	return shipment, nil
}
