package entity

type Car struct {
	ID        int64   `gorm:"primaryKey;AUTO_INCREMENT" json:"id"`
	CarName   string  `gorm:"column(car_name)" json:"car_name"`
	DayRate   float64 `gorm:"column(day_rate)" json:"day_rate"`
	MonthRate float64 `gorm:"column(month_rate)" json:"month_rate"`
	Image     string  `gorm:"column(image)" json:"image"`
	IsDeleted int     `gorm:"DEFAULT:1" json:"-"`
}

func (m *Car) TableName() string {
	return "cars"
}
