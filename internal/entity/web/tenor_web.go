package web

type CreateTenorRequest struct {
	Data []uint8 `json:"data" validate:"required"`
}
