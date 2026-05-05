package po

type Role struct {
	ID   int64  `gorm:"colum:id;primaryKey;autoIncrement" json:"id"`
	Name string `gorm:"column:name;type:varchar(50);unique;not null" json:"name"`
	Note string `gorm:"column:note;type:text" json:"note"`
	IsActive bool `gorm:"column:is_active;type:tinyint(1);default:1" json:"is_active"`
}

func (r *Role) TableName() string {
	return "roles"
}
