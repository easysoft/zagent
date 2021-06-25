package model

type Container struct {
	BaseModel

	Ident string
}

func (Container) TableName() string {
	return "biz_container"
}
