package agentModel

import (
	"time"
)

type BaseModel struct {
	//gorm.Model

	ID        uint       `gorm:"primary_key" sql:"type:INT(10) UNSIGNED NOT NULL" json:"id"`
	CreatedAt *time.Time `json:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
	DeletedAt *time.Time `sql:"index" json:"deletedAt,omitempty"`

	Deleted  bool `json:"deleted,omitempty" gorm:"default:false"`
	Disabled bool `json:"disabled,omitempty" gorm:"default:false"`
}

var (
	Models = []interface{}{
		&Task{},
	}
)
