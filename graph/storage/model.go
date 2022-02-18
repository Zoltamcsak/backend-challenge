package storage

type Country string

const (
	ITALY  Country = "ITALY"
	FRANCE Country = "FRANCE"
)

type Salary struct {
	ID          uint
	Gross       float64
	Country     Country
	Bonus       float64
	Month       int
	Year        int
	UserID      uint
	UserProfile UserProfile `gorm:"foreignKey:UserID"`
}

type UserProfile struct {
	ID                uint
	FirstName         string
	LastName          string
	ProfilePictureUrl string
	Salaries          []Salary `gorm:"foreignKey:UserID"`
}
