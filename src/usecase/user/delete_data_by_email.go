package users

type DeleteDataByEmailRequest struct {
	Email string `json:"email" query:"email" form:"email"`
}
