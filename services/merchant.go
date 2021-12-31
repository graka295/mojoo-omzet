package services

import (
	"mojoo/omzet/container"
	"mojoo/omzet/model/domain"
	"mojoo/omzet/model/dto"
	"time"

	"github.com/labstack/echo/v4"
)

// Merchant contract
type Merchant interface {
	GetOmzetMerchant(c echo.Context, userID int) error
}

// MerchantImpl struct Merchant
type MerchantImpl struct {
	container container.Container
}

func NewMerchant(container container.Container) Merchant {
	return MerchantImpl{
		container: container,
	}
}
func (x MerchantImpl) GetOmzetMerchant(c echo.Context, userID int) error {
	req := dto.OmzetMerchant{}
	if err := c.Bind(&req); err != nil {
		return x.container.GetResponse().InternalServerError(c, err.Error())
	}
	messageValidate, valid := x.container.GetValidate().Validate(req)
	if !valid {
		return x.container.GetResponse().BadRequest(c, messageValidate)
	}
	merchant := domain.NewMerchants()
	data := merchant.FindbyID(x.container, req.ID)
	if data == nil {
		return x.container.GetResponse().Notfound(c, "merchant not found")
	}
	if userID != data.UserID {
		return x.container.GetResponse().Forbidden(c, "your dont have access for this merchant")
	}
	t := time.Date(req.Year, time.Month(req.Month), 0, 0, 0, 0, 0, time.UTC)
	totalDay := t.Day()
	var totalPage int
	var limitPage int
	if req.Page == 1 {
		totalPage = 15
		limitPage = 1
	} else {
		totalPage = req.Page * 15
		limitPage = (req.Page - 1) * 15
	}
	res := []dto.ResOmzetMerchant{}
	for i := 1; limitPage <= totalDay; i++ {
		if i > totalPage {
			break
		}
		time := time.Date(req.Year, time.Month(req.Month), i, 0, 0, 0, 0, time.UTC)
		sum := merchant.Sum(x.container, req.ID, time)
		res = append(res, dto.ResOmzetMerchant{
			Date:  time.Format("2006-01-02"),
			Omzet: sum,
		})
	}
	return x.container.GetResponse().Ok(c, res)
}
