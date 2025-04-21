package shipment

import (
	repository "CKit/internal/repository/shipment"
	"context"
)

// DeleteShipmentUseCase handles deleting a shipment.
type DeleteShipmentUseCase struct {
	repo repository.ShipmentRepository
}

// NewDeleteShipmentUseCase initializes the use case.
func NewDeleteShipmentUseCase(r repository.ShipmentRepository) *DeleteShipmentUseCase {
	return &DeleteShipmentUseCase{repo: r}
}

// Execute deletes a shipment by ID.
func (uc *DeleteShipmentUseCase) Execute(ctx context.Context, id string) error {
	return uc.repo.Delete(ctx, id)
}
