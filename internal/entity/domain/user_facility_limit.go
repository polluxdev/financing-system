package domain

type UserFacilityLimit struct {
	ID          string  `db:"id"`
	UserID      string  `db:"user_id"`
	LimitAmount float64 `db:"limit_amount"`
}

func (UserFacilityLimit) TableName() string {
	return "user_facility_limits"
}
