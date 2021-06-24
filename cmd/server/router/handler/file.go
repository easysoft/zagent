package handler

import (
	_const "github.com/easysoft/zagent/internal/pkg/const"
	_dateUtils "github.com/easysoft/zagent/internal/pkg/lib/date"
	_fileUtils "github.com/easysoft/zagent/internal/pkg/lib/file"
	uuid "github.com/iris-contrib/go.uuid"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"mime/multipart"
	"time"
)

type FileCtrl struct {
	Ctx iris.Context
}

func NewFileCtrl() *FileCtrl {
	return &FileCtrl{}
}
func (g *FileCtrl) PostUpload() {
	dir := _const.UploadDir + _dateUtils.DateStr(time.Now())
	_fileUtils.MkDirIfNeeded(dir)

	g.Ctx.UploadFormFiles("./uploads", beforeFileSave)
}

func beforeFileSave(context *context.Context, file *multipart.FileHeader) bool {
	uuid, _ := uuid.NewV4()
	file.Filename = uuid.String() + "-" + file.Filename

	return true
}
