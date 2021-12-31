package services

import (
	"fmt"
	"mojoo/omzet/container"
	"mojoo/omzet/model/domain"
	"mojoo/omzet/model/dto"
	"time"

	"github.com/labstack/echo/v4"
)

// OutLets contract
type OutLets interface {
	GetOmszetOutLets(c echo.Context, userID int) error
}

// OutLetsImpl struct OutLets
type OutLetsImpl struct {
	container container.Container
}

func NewOustLets(container container.Container) OutLets {
	return OutLetsImpl{
		container: container,
	}
}
func (x OutLetsImpl) GetOmszetOutLets(c echo.Context, userID int) error {
	req := dto.OmzetOutlets{}
	if err := c.Bind(&req); err != nil {
		return x.container.GetResponse().InternalServerError(c, err.Error())
	}
	messageValidate, valid := x.container.GetValidate().Validate(req)
	if !valid {
		return x.container.GetResponse().BadRequest(c, messageValidate)
	}
	OutLets := domain.NewOutlets()
	data := OutLets.FindbyID(x.container, req.ID)
	if data == nil {
		return x.container.GetResponse().Notfound(c, "OutLets not found")
	}
	fmt.Println(userID, data.UserID)
	if userID != data.UserID {
		return x.container.GetResponse().Forbidden(c, "your dont have access for this OutLets")
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
	res := []dto.ResOmzetOutlets{}
	for i := 1; limitPage <= totalDay; i++ {
		if i > totalPage {
			break
		}
		time := time.Date(req.Year, time.Month(req.Month), i, 0, 0, 0, 0, time.UTC)
		sum := OutLets.Sum(x.container, req.ID, time)
		res = append(res, dto.ResOmzetOutlets{
			Date:  time.Format("2006-01-02"),
			Omzet: sum,
		})
	}
	return x.container.GetResponse().Ok(c, res)
}
