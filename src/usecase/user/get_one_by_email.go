package users

type GetOneByEmailRequest struct {
	Email string `json:"email" query:"email" form:"email"`
}
