package db

import (
	agentConf "github.com/easysoft/zagent/internal/agent/conf"
	agentConst "github.com/easysoft/zagent/internal/agent/utils/const"
	_fileUtils "github.com/easysoft/zagent/internal/pkg/libs/file"
	_logUtils "github.com/easysoft/zagent/internal/pkg/libs/log"
	"gorm.io/driver/sqlite"
	_ "gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"gorm.io/plugin/dbresolver"
	"path/filepath"
	"strings"
	"time"
)

var (
	inst *Instance

	FlagVarDBFile string
)

func GetInst() *Instance {
	if inst == nil {
		InitDB()
	}

	return inst
}

func InitDB() {
	var dialector gorm.Dialector

	conn := DBFile()
	dialector = sqlite.Open(conn)

	DB, err := gorm.Open(dialector, &gorm.Config{
		SkipDefaultTransaction: false,
		Logger:                 logger.Default.LogMode(logger.Info),
		NamingStrategy: schema.NamingStrategy{
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
	inst.config = &agentConf.Inst
}

func (*Instance) DB() *gorm.DB {
	return inst.db
}

type Instance struct {
	config *agentConf.Config
	db     *gorm.DB
}

func (i *Instance) Close() error {
	if i.db != nil {
		sqlDB, _ := i.db.DB()
		return sqlDB.Close()
	}
	return nil
}

func DBFile() string {
	path := filepath.Join(_fileUtils.GetExeDir(), strings.ToLower(agentConst.AppName+".db"))
	return path
}
