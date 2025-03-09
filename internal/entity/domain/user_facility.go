package domain

import "time"

type UserFacility struct {
	ID                  string    `db:"id"`
	UserID              string    `db:"user_id"`
	UserFacilityLimitID string    `db:"user_facility_limit_id"`
	Amount              float64   `db:"amount"`
	Tenor               uint8     `db:"tenor"`
	StartDate           time.Time `db:"start_date"`
	MonthlyInstallment  float64   `db:"monthly_installment"`
	TotalMargin         float64   `db:"total_margin"`
	TotalPayment        float64   `db:"total_payment"`
}

func (UserFacility) TableName() string {
	return "user_facilities"
}
