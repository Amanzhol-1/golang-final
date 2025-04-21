package shipment

import (
	"CKit/internal/entity"
	repository "CKit/internal/repository/shipment"
	"context"

	"github.com/google/uuid"
)

// CreateShipmentUseCase handles creating shipments.
type CreateShipmentUseCase struct {
	repo repository.ShipmentRepository
}

// NewCreateShipmentUseCase initializes the use case.
func NewCreateShipmentUseCase(r repository.ShipmentRepository) *CreateShipmentUseCase {
	return &CreateShipmentUseCase{repo: r}
}

// Execute creates a new shipment with generated ID.
func (uc *CreateShipmentUseCase) Execute(ctx context.Context, shipment *entity.Shipment) (*entity.Shipment, error) {
	shipment.ID = uuid.New().String()
	if err := uc.repo.Save(ctx, shipment); err != nil {
		return nil, err
	}
	return shipment, nil
}
