package model

import (
	"github.com/JumpSama/aug-blog/pkg/auth"
	"github.com/JumpSama/aug-blog/pkg/constvar"
	"github.com/go-playground/validator"
	"github.com/jinzhu/gorm"
	"sync"
)

// 用户表
type User struct {
	gorm.Model
	Account  string `gorm:"size:20;not null;unique" json:"account" binding:"required" validate:"min=5,max=20"`
	Username string `gorm:"size:20;not null" json:"username" binding:"required" validate:"min=1,max=20"`
	Password string `gorm:"not null" json:"password" binding:"required" validate:"min=5,max=255"`
	Status   bool   `gorm:"not null;default:true" json:"status"`
}

// 列表请求字段
type ListRequest struct {
	Account  string `json:"account"`
	Username string `json:"username"`
	Offset   int    `json:"offset"`
	Limit    int    `json:"limit"`
}

// 用户列表字段
type UserInfo struct {
	Id        uint   `json:"id"`
	Account   string `json:"account"`
	Username  string `json:"username"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

// 用户列表
type UserList struct {
	Lock  *sync.Mutex
	IdMap map[uint]*UserInfo
}

// 创建用户
func (u *User) Create() error {
	return DB.Self.Create(&u).Error
}

// 删除用户
func DeleteUser(id uint) error {
	self := User{}
	self.ID = id
	return DB.Self.Delete(&self).Error
}

// 更新用户
func (u *User) Update() error {
	return DB.Self.Save(&u).Error
}

// 通过账号获取用户
func GetUserByAccount(account string) (*User, error) {
	u := &User{}
	d := DB.Self.Where("account = ?", account).First(&u)
	return u, d.Error
}

// 通过账号(和id)获取用户数量
func GetUserCountByAccount(account string, id uint) int {
	count := 0

	sql := DB.Self.Where("account = ?", account)

	if id > 0 {
		sql = sql.Where("id <> ?", id)
	}

	sql.Model(&User{}).Count(&count)

	return count
}

// 用户列表
func GetUserList(req *ListRequest) (list []*User, count int) {
	if req.Limit == 0 {
		req.Limit = constvar.DefaultLimit
	}

	count = 0
	list = make([]*User, 0)

	sql := DB.Self

	if req.Account != "" {
		sql = sql.Where("account like ?", "%"+req.Account+"%")
	}

	if req.Username != "" {
		sql = sql.Where("username like ?", "%"+req.Username+"%")
	}

	sql.Model(&User{}).Count(&count).Offset(req.Offset).Limit(req.Limit).Find(&list)

	return
}

// 密码验证
func (u *User) Compare(pwd string) bool {
	return auth.Compare(u.Password, pwd)
}

// 密码加密
func (u *User) Encrypt() (err error) {
	u.Password, err = auth.Encrypt(u.Password)
	return
}

// 字段验证
func (u *User) Validate() error {
	validate := validator.New()
	return validate.Struct(u)
}
