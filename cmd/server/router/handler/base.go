package handler

import (
	_const "github.com/easysoft/zagent/internal/pkg/const"
	_fileUtils "github.com/easysoft/zagent/internal/pkg/lib/file"
	_httpUtils "github.com/easysoft/zagent/internal/pkg/lib/http"
	"github.com/easysoft/zagent/internal/server/biz/validate"
	"github.com/go-playground/validator/v10"
	"github.com/gofrs/uuid"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"mime/multipart"
	"strings"
)

type BaseCtrl struct {
	Ctx iris.Context
}

func (c *BaseCtrl) Validate(s interface{}, ctx iris.Context) bool {
	err := validate.Validate.Struct(s)

	if err != nil {
		errs := err.(validator.ValidationErrors)
		for _, e := range errs.Translate(validate.ValidateTrans) {
			if len(e) > 0 {
				_, _ = ctx.JSON(_httpUtils.ApiRes(400, e, nil))
				return true
			}
		}
	}

	return false
}

func beforeFileSave(context *context.Context, file *multipart.FileHeader) bool {
	name := strings.ReplaceAll(file.Filename, "\\", _const.PthSep)

	uuid, _ := uuid.NewV4()
	file.Filename = _fileUtils.GetFileNameWithoutExt(name) +
		"-" + uuid.String() +
		_fileUtils.GetExtName(name)

	return true
}
