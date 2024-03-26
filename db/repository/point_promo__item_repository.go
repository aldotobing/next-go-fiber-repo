package repository

import (
	"context"
	"database/sql"

	"nextbasis-service-v-0.1/usecase/viewmodel"
)

// IPointPromoItemRepository ...
type IPointPromoItemRepository interface {
	AddBulk(c context.Context, promoID string, items []viewmodel.PointPromoItemVM) error
	Delete(c context.Context, promoID string) error
}

// PointPromoItemRepository ...
type PointPromoItemRepository struct {
	DB *sql.DB
}

// NewPointPromoItemRepository ...
func NewPointPromoItemRepository(DB *sql.DB) IPointPromoItemRepository {
	return &PointPromoItemRepository{DB: DB}
}

// AddBulk ...
func (repository PointPromoItemRepository) AddBulk(c context.Context, promoID string, items []viewmodel.PointPromoItemVM) (err error) {
	var statementInsert string

	for _, datum := range items {
		if statementInsert == "" {
			statementInsert += `('` + promoID + `', '` + datum.ID + `', '` + datum.UomID + `', '` + datum.UomName + `', '` + datum.Convertion + `', NOW(), NOW())`
		} else {
			statementInsert += `,('` + promoID + `', '` + datum.ID + `', '` + datum.UomID + `', '` + datum.UomName + `','` + datum.Convertion + `', NOW(), NOW())`
		}
	}

	statement := `INSERT INTO POINT_PROMO_ITEM (
			PROMO_ID, 
			ITEM_ID,
			UOM_ID,
			UOM_NAME,
			CONVERTION,
			CREATED_AT,
			UPDATED_AT
		)
	VALUES ` + statementInsert
	err = repository.DB.QueryRowContext(c, statement).Err()

	return
}

// Delete ...
func (repository PointPromoItemRepository) Delete(c context.Context, promoID string) (err error) {
	statement := `UPDATE POINT_PROMO_ITEM SET 
	DELETED_AT = NOW()
	WHERE PROMO_ID = ` + promoID

	err = repository.DB.QueryRowContext(c, statement).Err()

	return
}
