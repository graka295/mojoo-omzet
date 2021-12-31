package dto

type (
	// LoginRequest struct for request login
	LoginRequest struct {
		Username string `validate:"required,min=1,max=45" json:"username"`
		Password string `validate:"required,max=40" json:"password"`
	}
	// LoginResponse struct for Response login
	LoginResponse struct {
		Name     string `json:"name"`
		Username string `json:"username"`
		Token    string `json:"token"`
	}
)
