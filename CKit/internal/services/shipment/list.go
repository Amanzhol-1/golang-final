package shipment

import (
	"CKit/internal/entity"
	repository "CKit/internal/repository/shipment"
	"context"
)

// ListShipmentsUseCase handles fetching all shipments.
type ListShipmentsUseCase struct {
	repo repository.ShipmentRepository
}

// NewListShipmentsUseCase initializes the use case.
func NewListShipmentsUseCase(r repository.ShipmentRepository) *ListShipmentsUseCase {
	return &ListShipmentsUseCase{repo: r}
}

// Execute returns all shipments.
func (uc *ListShipmentsUseCase) Execute(ctx context.Context) ([]*entity.Shipment, error) {
	return uc.repo.List(ctx)
}
