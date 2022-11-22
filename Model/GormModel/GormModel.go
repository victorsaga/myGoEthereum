package GormModel

type Account struct {
	Name     string `gorm:"column:name;primary_key"`
	Password string `gorm:"column:password;NOT NULL"`
	Enable   int8   `gorm:"column:enable;NOT NULL"`
}
