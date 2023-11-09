package _db

import (
	"fmt"
	agentConf "github.com/easysoft/zagent/internal/pkg/conf"
	consts "github.com/easysoft/zagent/internal/pkg/const"
	_fileUtils "github.com/easysoft/zagent/pkg/lib/file"
	_logUtils "github.com/easysoft/zagent/pkg/lib/log"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"gorm.io/plugin/dbresolver"
	"path/filepath"
	"time"

	_ "github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

var (
	inst *Instance

	FlagVarDBFile string
)

func GetInst() *Instance {
	return inst
}

func InitDB(mode string) {
	var dialector gorm.Dialector

	if mode == "host" {
		conn := DBFile(mode)
		dialector = sqlite.Open(conn)
	}

	prefix := agentConf.Inst.DB.Prefix

	DB, err := gorm.Open(dialector, &gorm.Config{
		SkipDefaultTransaction: false,
		Logger:                 logger.Default.LogMode(logger.Info),
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   prefix,
			SingularTable: false,
		},
	})

	if err != nil {
		_logUtils.Info(err.Error())
	}

	_ = DB.Use(
		dbresolver.Register(
			dbresolver.Config{ /* xxx */ }).
			SetConnMaxIdleTime(time.Hour).
			SetConnMaxLifetime(24 * time.Hour).
			SetMaxIdleConns(100).
			SetMaxOpenConns(200),
	)

	DB.Session(&gorm.Session{FullSaveAssociations: true, AllowGlobalUpdate: false})

	inst = &Instance{}
	inst.db = DB
}

func (*Instance) DB() *gorm.DB {
	return inst.db
}

type Instance struct {
	//config *serverConf.DBConfig
	db *gorm.DB
}

func (i *Instance) Close() error {
	if i.db != nil {
		sqlDB, _ := i.db.DB()
		return sqlDB.Close()
	}
	return nil
}

func DBFile(mode string) string {
	path := filepath.Join(_fileUtils.GetExeDir(), fmt.Sprintf("%s-%s.db", consts.AppName, mode))
	return path
}
