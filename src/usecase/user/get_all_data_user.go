package users

type GetAllDataUsersRequest struct {
	Keywords string `json:"keywords" query:"keywords" form:"keywords"`
}
