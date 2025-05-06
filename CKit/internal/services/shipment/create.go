package shipment

import (
	"CKit/internal/entity"
	"CKit/internal/middleware"
	repository "CKit/internal/repository/shipment"
	"context"
	"errors"

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
	subInfo *middleware.SubscriptionInfo,
	in *entity.CreateShipmentRequest,
) (*entity.Shipment, error) {
	if subInfo == nil || subInfo.IsActive == false {
		return nil, errors.New("ErrNoActiveSubscription")
	}
	if in.PickupTime.Before(subInfo.StartDate) || in.PickupTime.After(subInfo.EndDate) {
		return nil, errors.New("ErrPickupOutsideSubscription")
	}
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
		Status:          entity.StatusPending,
	}
	if err := uc.repo.Save(ctx, s); err != nil {
		return nil, err
	}
	return s, nil
}
