package web

type CreateUserRequest struct {
	Name        string `json:"name" validate:"required"`
	PhoneNumber string `json:"phone_number" validate:"required,isValidPhoneNumber"`
}
