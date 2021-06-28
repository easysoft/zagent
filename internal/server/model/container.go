package model

type Container struct {
	BaseModel

	Ident string `json:"ident"`
}

func (Container) TableName() string {
	return "biz_container"
}
