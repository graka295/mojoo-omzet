package domain

import (
	"errors"
	"mojoo/omzet/container"
	"time"

	"gorm.io/gorm"
)

type Outlets interface {
	FindbyID(ct container.Container, id int) *OutletsImpl
	Sum(ct container.Container, idMerchant int, date time.Time) int
}
type OutletsImpl struct {
	ID           int    `gorm:"column:id"`
	UserID       int    `gorm:"column:user_id"`
	MerchantName string `gorm:"column:merchant_name"`
}

func NewOutlets() Outlets {
	return &OutletsImpl{}
}
func (x OutletsImpl) TableName() string {
	return "Outlets"
}

func (x OutletsImpl) FindbyID(ct container.Container, id int) *OutletsImpl {
	var data OutletsImpl
	res := ct.GetConnection().Select("Merchants.id as id,Merchants.user_id as user_id,Merchants.merchant_name").Table("Outlets").Joins("JOIN Merchants ON Merchants.id = Outlets.merchant_id").Where("Outlets.id = ?", id).First(&data)
	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return nil
	}

	if res.Error != nil {
		ct.GetLoger().LogError(res.Error)
		return nil
	}
	return &data
}

func (x OutletsImpl) Sum(ct container.Container, idOutlets int, date time.Time) int {
	var sum int
	dates := date.Format("2006-01-02")
	ct.GetConnection().Table("Transactions").Select("SUM(bill_total) as omzet").Where("DATE(Transactions.created_at) = ? and Transactions.outlet_id =? ", dates, idOutlets).Row().Scan(&sum)
	return sum
}
