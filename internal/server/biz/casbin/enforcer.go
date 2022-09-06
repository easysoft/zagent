package bizCasbin

import (
	"fmt"
	"github.com/casbin/casbin/v2"
	_db "github.com/easysoft/zv/pkg/db"
	_commonUtils "github.com/easysoft/zv/pkg/lib/common"
	_fileUtils "github.com/easysoft/zv/pkg/lib/file"
	_logUtils "github.com/easysoft/zv/pkg/lib/log"
	"github.com/easysoft/zv/res/server"
	"github.com/sirupsen/logrus"
	"path/filepath"
)

func NewEnforcer() *casbin.Enforcer {
	adapter, err := NewAdapterByDB(_db.GetInst().DB())
	if err != nil {
		logrus.Println(fmt.Sprintf("NewAdapter 错误: %v", err))
	}

	exeDir := _fileUtils.GetExeDir()
	pth := ""
	enforcer := &casbin.Enforcer{}
	if _commonUtils.IsRelease() {
		pth = filepath.Join(exeDir, "rbac_model.conf")
		if !_fileUtils.FileExist(pth) {
			bytes, _ := serverRes.Asset("res/server/rbac_model.conf")
			_fileUtils.WriteFile(pth, string(bytes))
		}
	} else {
		pth = filepath.Join(exeDir, "cmd", "server", "rbac_model.conf")
	}

	_logUtils.Infof("从文件%s加载casbin配置", pth)
	enforcer, err = casbin.NewEnforcer(pth, adapter)
	if err != nil {
		logrus.Println(fmt.Sprintf("NewEnforcer 错误: %v", err))
	}

	_ = enforcer.LoadPolicy()

	return enforcer
}
