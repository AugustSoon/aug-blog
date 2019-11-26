package user

import "github.com/JumpSama/aug-blog/model"

type CreateResponse struct {
	ID       uint   `json:"id"`
	Account  string `json:"account"`
	Username string `json:"username"`
}

type ListResponse struct {
	Total int               `json:"total"`
	List  []*model.UserInfo `json:"list"`
}
