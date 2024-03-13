package request

import (
	"fmt"

	"github.com/naufaldinta13/cars/config"
	"github.com/naufaldinta13/cars/entity"
)

type GetRequest struct {
	Search string `json:"search"`
	Limit  int64  `json:"limit"`
	Page   int64  `json:"page"`
}

func (r *GetRequest) GetLimit() int64 {
	if r.Limit == 0 {
		r.Limit = 1
	}

	return r.Limit
}

func (r *GetRequest) GetOffset() int64 {
	if r.Page == 0 {
		r.Page = 1
	}

	return r.GetLimit() * (r.Page - 1)

}

func (r *GetRequest) Get() (data *config.ResponseBody) {
	var mx []*entity.Car
	q := config.GetDB().Model(&mx)

	q = q.Where("is_deleted = 0")

	if r.Search != "" {
		q = q.Where("car_name LIKE ?", fmt.Sprintf("%s%s%s", "%", r.Search, "%"))
	}

	var total int64
	e := q.Count(&total).Error
	if e != nil || total == 0 {
		return
	}

	q.Limit(r.GetLimit())
	q.Offset(r.GetOffset())

	if e = q.Find(&mx).Error; e == nil {
		data = &config.ResponseBody{
			Success: true,
			Total:   total,
			Data:    mx,
		}
	}

	return
}

func (r *GetRequest) Show(id string) (data *config.ResponseBody) {

	var mx entity.Car
	if e := config.GetDB().Model(&mx).Where("id = ?", id).First(&mx).Error; e == nil {
		data = &config.ResponseBody{
			Success: true,
			Data:    mx,
		}
	}

	return
}
