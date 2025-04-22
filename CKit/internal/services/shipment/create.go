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

func (uc *CreateShipmentUseCase) Execute(
	ctx context.Context,
	userID string,
	in *entity.CreateShipmentRequest,
) (*entity.Shipment, error) {
	s := &entity.Shipment{
		ID:              uuid.NewString(),
		UserID:          userID,
		FromAddress:     in.FromAddress,
		ToAddress:       in.ToAddress,
		PickupTime:      in.PickupTime,
		DeliveryPrice:   in.DeliveryPrice,
		PriceNegotiable: in.PriceNegotiable,
		Weight:          in.Weight,
		Volume:          in.Volume,
		CargoType:       in.CargoType,
		SenderName:      in.SenderName,
		SenderPhone:     in.SenderPhone,
		ReceiverName:    in.ReceiverName,
		ReceiverPhone:   in.ReceiverPhone,
		AdditionalNotes: in.AdditionalNotes,
	}
	if err := uc.repo.Save(ctx, s); err != nil {
		return nil, err
	}
	return s, nil
}
