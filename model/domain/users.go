package domain

import (
	"errors"
	"mojoo/omzet/container"
	"time"

	"gorm.io/gorm"
)

type Users interface {
	Login(ct container.Container, username, password string) *UsersImpl
}
type UsersImpl struct {
	ID        int       `gorm:"column:id"`
	Name      string    `gorm:"column:name"`
	Username  string    `gorm:"column:user_name"`
	Password  string    `gorm:"column:password"`
	CreatedAt time.Time `gorm:"column:created_at"`
	CreatedBy int       `gorm:"column:created_by"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
	UpdatedBy int       `gorm:"column:updated_by"`
}

func NewUsers() Users {
	return &UsersImpl{}
}
func (x UsersImpl) TableName() string {
	return "Users"
}

func (x UsersImpl) Login(ct container.Container, username, password string) *UsersImpl {
	var data UsersImpl
	res := ct.GetConnection().Where("user_name = ? and password = md5(?)", username, password).First(&data)
	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return nil
	}

	if res.Error != nil {
		ct.GetLoger().LogError(res.Error)
		return nil
	}
	return &data
}
