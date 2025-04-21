package repository

import (
	"CKit/internal/entity"
	"context"

	"gorm.io/gorm"
)

// ShipmentRepository defines CRUD operations for shipments.
type ShipmentRepository interface {
	// Save создает новую заявку на доставку
	Save(ctx context.Context, shipment *entity.Shipment) error
	// FindByID возвращает заявку по её ID
	FindByID(ctx context.Context, id string) (*entity.Shipment, error)
	// Update обновляет существующую заявку
	Update(ctx context.Context, shipment *entity.Shipment) error
	// Delete удаляет заявку по ID
	Delete(ctx context.Context, id string) error
	// List возвращает все заявки
	List(ctx context.Context) ([]*entity.Shipment, error)
}

type DBShipmentRepository struct {
	db *gorm.DB
}

// NewGormShipmentRepository creates a new GORM-based repository.
func NewDBShipmentRepository(db *gorm.DB) *DBShipmentRepository {
	return &DBShipmentRepository{db: db}
}

// Save creates a new shipment record.
func (r *DBShipmentRepository) Save(ctx context.Context, shipment *entity.Shipment) error {
	return r.db.WithContext(ctx).Create(shipment).Error
}

// FindByID retrieves a shipment by its ID.
func (r *DBShipmentRepository) FindByID(ctx context.Context, id string) (*entity.Shipment, error) {
	var s entity.Shipment
	if err := r.db.WithContext(ctx).First(&s, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &s, nil
}

// Update modifies an existing shipment record.
func (r *DBShipmentRepository) Update(ctx context.Context, shipment *entity.Shipment) error {
	return r.db.WithContext(ctx).Save(shipment).Error
}

// Delete removes a shipment by its ID.
func (r *DBShipmentRepository) Delete(ctx context.Context, id string) error {
	return r.db.WithContext(ctx).Delete(&entity.Shipment{}, "id = ?", id).Error
}

// List returns all shipments.
func (r *DBShipmentRepository) List(ctx context.Context) ([]*entity.Shipment, error) {
	var list []*entity.Shipment
	if err := r.db.WithContext(ctx).Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}
