package model

type UserInfo struct {
	Id        uint   `json:"id"`
	Username  string `json:"username"`
	SayHello  string `json:"sayHello"`
	Password  string `json:"password"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

//type UserList struct {
//	Lock  *sync.Mutex
//	IdMap map[uint]*UserInfo
//}

type Token struct {
	Token string `json:"token"`
}
