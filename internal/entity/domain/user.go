package domain

type User struct {
	ID          string `db:"id"`
	Name        string `db:"name"`
	PhoneNumber string `db:"phone_number"`
}

func (User) TableName() string {
	return "users"
}
