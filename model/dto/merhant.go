package dto

type (
	// OmzetMerchant struct for Response login
	OmzetMerchant struct {
		ID    int `json:"id" validate:"required"`
		Month int `json:"month" validate:"required"`
		Page  int `json:"page"`
		Year  int `json:"year" validate:"required"`
	}
	// ResOmzetMerchant struct for Response login
	ResOmzetMerchant struct {
		Date  string `json:"date" `
		Omzet int    `json:"omzet"`
	}
)
