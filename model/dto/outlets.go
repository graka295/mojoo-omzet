package dto

type (
	// OmzetOutlets struct for Response login
	OmzetOutlets struct {
		ID    int `json:"id" validate:"required"`
		Month int `json:"month" validate:"required"`
		Page  int `json:"page"`
		Year  int `json:"year" validate:"required"`
	}
	// ResOmzetOutlets struct for Response login
	ResOmzetOutlets struct {
		Date  string `json:"date" `
		Omzet int    `json:"omzet"`
	}
)
