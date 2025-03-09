package domain

import "time"

type UserFacilityDetail struct {
	ID                string    `db:"id"`
	UserFacilityID    string    `db:"user_facility_id"`
	DueDate           time.Time `db:"due_date"`
	InstallmentAmount float64   `db:"installment_amount"`
}

func (UserFacilityDetail) TableName() string {
	return "user_facility_details"
}
