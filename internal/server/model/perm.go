package model

type Permission struct {
	BaseModel

	Name        string `gorm:"not null;type:varchar(256)" json:"name" validate:"required,gte=4,lte=50" comment:"名称"`
	DisplayName string `gorm:"type:varchar(256)" json:"displayName" comment:"显示名称"`
	Description string `gorm:"type:varchar(256)" json:"description" comment:"描述"`
	Act         string `gorm:"type:varchar(256)" json:"act" comment:"动作"`
}

func (Permission) TableName() string {
	return "biz_permission"
}
