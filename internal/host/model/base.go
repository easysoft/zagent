package agentModel

import (
	"time"
)

type BaseModel struct {
	//gorm.Model

	ID          uint       `json:"id" gorm:"primary_key" sql:"type:INT(10) UNSIGNED NOT NULL"`
	CreatedDate *time.Time `json:"createdDate" gorm:"autoCreateTime;column:createdDate"`
	UpdatedDate *time.Time `json:"updatedDate,omitempty" gorm:"autoUpdateTime;column:updatedDate"`
	DeletedDate *time.Time `json:"deletedDate,omitempty" gorm:"autoDeleteTime;column:deletedDate"`

	Deleted  bool `json:"deleted,omitempty" gorm:"default:false"`
	Disabled bool `json:"disabled,omitempty" gorm:"default:false"`
}

var (
	Models = []interface{}{
		&Task{},
	}
)
