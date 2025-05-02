package account

type CreateAccountRequest struct {
	PhoneNumber string `json:"number" binding:"required"`
	Owner       string `json:"owner" binding:"required"`
	Currency    string `json:"currency" binding:"required,oneof=USD EUR"`
}

type UpdateAccountRequest struct {
	PhoneNumber string `json:"number" binding:"required"`
	Owner       string `json:"owner" binding:"required"`
	Currency    string `json:"currency" binding:"required,oneof=USD EUR"`
	Email       string `json:"email"`
	Balance     int64  `json:"balance" binding:"required, min=1"`
}

type AccountResponse struct {
	ID          int64  `json:"id"`
	Email       string `json:"email"`
	PhoneNumber string `json:"number"`
	Owner       string `json:"owner"`
	Currency    string `json:"currency"`
	Balance     int64  `json:"balance"`
}

type AccountUriRequest struct {
	ID int64 `uri:"id" binding:"required"`
}

type GetAccountListRequest struct {
	PageId   int32 `form:"page" binding:"required,min=1"`
	PageSize int32 `form:"size" binding:"required,min=5,max=10"`
}
