package handler

import (
	"fmt"
	consts "github.com/easysoft/zv/internal/pkg/const"
	"github.com/easysoft/zv/internal/server/biz/validate"
	"github.com/easysoft/zv/internal/server/model"
	serverService "github.com/easysoft/zv/internal/server/service"
	_httpUtils "github.com/easysoft/zv/pkg/lib/http"
	"github.com/go-playground/validator/v10"
	"github.com/kataras/iris/v12"
)

type PermCtrl struct {
	PermService *serverService.PermService `inject:""`
}

func NewPermCtrl() *PermCtrl {
	return &PermCtrl{}
}

func (c *PermCtrl) GetPermission(ctx iris.Context) {
	id, _ := ctx.Params().GetUint("id")

	perm, err := c.PermService.GetPermission(id)
	if err != nil {
		_, _ = ctx.JSON(_httpUtils.RespData(consts.ResultPass, err.Error(), nil))
		return
	}

	_, _ = ctx.JSON(_httpUtils.RespData(consts.ResultPass, "操作成功", c.PermService.PermTransform(perm)))
}

func (c *PermCtrl) CreatePermission(ctx iris.Context) {
	perm := new(model.Permission)
	if err := ctx.ReadJSON(perm); err != nil {
		_, _ = ctx.JSON(_httpUtils.RespData(consts.ResultFail, err.Error(), nil))
		return
	}
	err := validate.Validate.Struct(*perm)
	if err != nil {
		errs := err.(validator.ValidationErrors)
		for _, e := range errs.Translate(validate.ValidateTrans) {
			if len(e) > 0 {
				_, _ = ctx.JSON(_httpUtils.RespData(consts.ResultFail, e, nil))
				return
			}
		}
	}

	err = c.PermService.CreatePermission(perm)
	if err != nil {
		_, _ = ctx.JSON(_httpUtils.RespData(consts.ResultFail, fmt.Sprintf("Error create prem: %s", err.Error()), nil))
		return
	}

	if perm.ID == 0 {
		_, _ = ctx.JSON(_httpUtils.RespData(consts.ResultFail, "操作失败", perm))
		return
	}
	_, _ = ctx.JSON(_httpUtils.RespData(consts.ResultPass, "操作成功", c.PermService.PermTransform(perm)))

}

func (c *PermCtrl) UpdatePermission(ctx iris.Context) {
	aul := new(model.Permission)

	if err := ctx.ReadJSON(aul); err != nil {
		_, _ = ctx.JSON(_httpUtils.RespData(consts.ResultFail, err.Error(), nil))
		return
	}
	err := validate.Validate.Struct(*aul)
	if err != nil {
		errs := err.(validator.ValidationErrors)
		for _, e := range errs.Translate(validate.ValidateTrans) {
			if len(e) > 0 {
				_, _ = ctx.JSON(_httpUtils.RespData(consts.ResultFail, e, nil))
				return
			}
		}
	}

	id, _ := ctx.Params().GetUint("id")
	err = c.PermService.UpdatePermission(id, aul)
	if err != nil {
		_, _ = ctx.JSON(_httpUtils.RespData(consts.ResultFail, fmt.Sprintf("Error update prem: %s", err.Error()), nil))
		return
	}

	_, _ = ctx.JSON(_httpUtils.RespData(consts.ResultPass, "操作成功", c.PermService.PermTransform(aul)))

}

func (c *PermCtrl) DeletePermission(ctx iris.Context) {
	id, _ := ctx.Params().GetUint("id")
	err := c.PermService.DeletePermissionById(id)
	if err != nil {
		_, _ = ctx.JSON(_httpUtils.RespData(consts.ResultFail, err.Error(), nil))
		return
	}
	_, _ = ctx.JSON(_httpUtils.RespData(consts.ResultPass, "删除成功", nil))
}

func (c *PermCtrl) GetAllPermissions(ctx iris.Context) {
	permissions, count, err := c.PermService.GetAllPermissions(nil)
	if err != nil {
		_, _ = ctx.JSON(_httpUtils.RespData(consts.ResultFail, err.Error(), nil))
		return
	}

	transform := c.PermService.PermsTransform(permissions)
	_, _ = ctx.JSON(_httpUtils.RespData(consts.ResultPass, "操作成功", map[string]interface{}{"items": transform, "total": count, "limit": "s.Limit"}))

}
