package model

import (
	"github.com/JumpSama/aug-blog/pkg/auth"
	"github.com/JumpSama/aug-blog/pkg/constvar"
	"github.com/go-playground/validator"
	"github.com/jinzhu/gorm"
)

type UserModel struct {
	gorm.Model
	Account  string `gorm:"size:20;not null;unique" json:"account" binding:"required" validate:"min=5,max=20"`
	Username string `gorm:"size:20;not null" json:"username" binding:"required" validate:"min=1,max=20"`
	Password string `gorm:"not null" json:"password" binding:"required" validate:"min=5,max=255"`
	Status   uint   `gorm:"not null;default:1"`
}

type ListRequest struct {
	Account  string `json:"account"`
	Username string `json:"username"`
	Offset   int    `json:"offset"`
	Limit    int    `json:"limit"`
}

func (u *UserModel) Create() error {
	return DB.Self.Create(&u).Error
}

func DeleteUser(id uint) error {
	self := UserModel{}
	self.ID = id
	return DB.Self.Delete(&self).Error
}

func (u *UserModel) Update() error {
	return DB.Self.Save(&u).Error
}

func GetUserByAccount(account string) (*UserModel, error) {
	u := &UserModel{}
	d := DB.Self.Where("account = ?", account).First(&u)
	return u, d.Error
}

func GetUserList(req *ListRequest) (list []*UserInfo, count int) {
	if req.Limit == 0 {
		req.Limit = constvar.DefaultLimit
	}

	count = 0
	list = make([]*UserInfo, 0)

	sql := DB.Self.Table("user_models")

	if req.Account != "" {
		sql = sql.Where("account like ?", "%"+req.Account+"%")
	}

	if req.Username != "" {
		sql = sql.Where("username like ?", "%"+req.Username+"%")
	}

	sql.Count(&count).Offset(req.Offset).Limit(req.Limit).Order("id desc").Find(&list)

	return
}

func (u *UserModel) Compare(pwd string) bool {
	return auth.Compare(u.Password, pwd)
}

func (u *UserModel) Encrypt() (err error) {
	u.Password, err = auth.Encrypt(u.Password)
	return
}

func (u *UserModel) Validate() error {
	validate := validator.New()
	return validate.Struct(u)
}
