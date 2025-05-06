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
func (uc *ListShipmentsUseCase) ExecuteCustomer(ctx context.Context, userId string) ([]*entity.Shipment, error) {
	return uc.repo.ListCustomer(ctx, userId)
}

func (uc *ListShipmentsUseCase) ExecuteDriver(ctx context.Context, userId string) ([]*entity.Shipment, error) {
	return uc.repo.ListDriver(ctx, userId)
}
