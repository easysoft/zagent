package agentModel

import (
	"time"
)

type BaseModel struct {
	//gorm.Model

	ID        uint       `gorm:"primary_key" sql:"type:INT(10) UNSIGNED NOT NULL" json:"id"`
	CreatedAt *time.Time `json:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt"`
	DeletedAt *time.Time `sql:"index" json:"deletedAt"`

	Deleted  bool `json:"deleted" gorm:"default:false"`
	Disabled bool `json:"disabled" gorm:"default:false"`
}

var (
	Models = []interface{}{
		&Task{},
	}
)
