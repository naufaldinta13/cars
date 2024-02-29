package entity

type Cars struct {
	CarID     int64   `orm:"column(car_id);auto" json:"car_id"`
	CarName   string  `orm:"column(car_name)" json:"car_name"`
	MonthRate float64 `orm:"column(month_rate)" json:"month_rate"`
	Image     string  `orm:"column(image)" json:"image"`
	IsDeleted bool    `orm:"column(is_deleted)" json:"-"`
}
