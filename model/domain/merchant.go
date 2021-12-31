package domain

import (
	"errors"
	"mojoo/omzet/container"
	"time"

	"gorm.io/gorm"
)

type Merchants interface {
	FindbyID(ct container.Container, id int) *MerchantsImpl
	Sum(ct container.Container, idMerchant int, date time.Time) int
}
type MerchantsImpl struct {
	ID           int       `gorm:"column:id"`
	UserID       int       `gorm:"column:user_id"`
	MerchantName string    `gorm:"column:merchant_name"`
	CreatedAt    time.Time `gorm:"column:created_at"`
	CreatedBy    int       `gorm:"column:created_by"`
	UpdatedAt    time.Time `gorm:"column:updated_at"`
	UpdatedBy    int       `gorm:"column:updated_by"`
}

func NewMerchants() Merchants {
	return &MerchantsImpl{}
}
func (x MerchantsImpl) TableName() string {
	return "Merchants"
}

func (x MerchantsImpl) FindbyID(ct container.Container, id int) *MerchantsImpl {
	var data MerchantsImpl
	res := ct.GetConnection().Where("id = ?", id).First(&data)
	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return nil
	}

	if res.Error != nil {
		ct.GetLoger().LogError(res.Error)
		return nil
	}
	return &data
}

func (x MerchantsImpl) Sum(ct container.Container, idMerchant int, date time.Time) int {
	var sum int
	dates := date.Format("2006-01-02")
	ct.GetConnection().Table("Transactions").Select("SUM(bill_total) as omzet").Where("DATE(Transactions.created_at) = ? and merchant_id = ? ", dates, idMerchant).Row().Scan(&sum)
	return sum
}
