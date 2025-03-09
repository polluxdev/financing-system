package web

import (
	"math"
)

type CalculateInstallmentRequest struct {
	Amount float64 `json:"amount" validate:"gt=0"`
}

type CalculateInstallmentDTO struct {
	CalculateInstallmentRequest
	Tenor              uint8   `json:"tenor"`
	TotalMargin        float64 `json:"total_margin"`
	TotalPayment       float64 `json:"total_payment"`
	MonthlyInstallment float64 `json:"monthly_installment"`
}

func ToCalculateInstallmentDTO(
	tenor uint8,
	amount, totalMargin, totalPayment, monthlyInstallment float64,
) CalculateInstallmentDTO {
	return CalculateInstallmentDTO{
		CalculateInstallmentRequest: CalculateInstallmentRequest{
			Amount: math.Round(amount*100) / 100,
		},
		Tenor:              tenor,
		TotalMargin:        math.Round(totalMargin*100) / 100,
		TotalPayment:       math.Round(totalPayment*100) / 100,
		MonthlyInstallment: math.Round(monthlyInstallment*100) / 100,
	}
}

type SubmitFinancingRequest struct {
	Amount    float64 `json:"amount" validate:"gt=0"`
	Tenor     uint8   `json:"tenor" validate:"gt=0"`
	StartDate string  `jon:"start_date" validate:"datetime=2006-01-02"`
}
