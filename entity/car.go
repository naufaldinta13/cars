package entity

type Cars struct {
	ID        int64   `orm:"column(id);auto" json:"id"`
	CarName   string  `orm:"column(car_name)" json:"car_name"`
	DayRate   float64 `orm:"column(day_rate)" json:"day_rate"`
	MonthRate float64 `orm:"column(month_rate)" json:"month_rate"`
	Image     string  `orm:"column(image)" json:"image"`
	IsDeleted int     `orm:"column(is_deleted)" json:"-"`
}
